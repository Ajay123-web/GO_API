package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func goDotEnvVariable(key string) string {
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
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