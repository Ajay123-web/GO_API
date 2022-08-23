package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Hackathon123:Team26@cluster0.ooyeu.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOpiton := options.Client().ApplyURI(connectionString)
	client , err := mongo.Connect(context.TODO() , clientOpiton)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo connection successful")
	collection = client.Database(dbName).Collection(colName)
}