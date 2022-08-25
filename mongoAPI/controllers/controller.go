package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	model "github.com/Ajay123-web/MONGODB/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func goDotEnvVariable(key string) string {
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatal("Error loading .env file")
  }

  return os.Getenv(key)
}
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	connectionString := goDotEnvVariable("MONGO_URI")
	clientOpiton := options.Client().ApplyURI(connectionString)
	client , err := mongo.Connect(context.TODO() , clientOpiton)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo connection successful")
	collection = client.Database(dbName).Collection(colName)
}

func insertOneMovie(movie model.Netflix) {
	inserted , err := collection.InsertOne(context.Background() , movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id , _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res , err := collection.UpdateOne(context.Background(), filter , update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", res.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id , _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount , err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got delete with delete count: ", deleteCount)
}

func getAllMovies() []primitive.M {
	cursor , err := collection.Find(context.Background() , bson.D{{}}) //we dont get data directly
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies , movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the API"))
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"]) 
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}