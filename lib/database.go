package lib

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var ReviewsCollection = connect()

func connect() *mongo.Collection {
	var ctx = context.Background()
	var mongoUrl = os.Getenv("MONGOURL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("bbpcontent")
	rc := database.Collection("reviews")
	return rc
}
