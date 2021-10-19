package reviewapi

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Review struct {
	ShopID            int    `bson:"shopid,omitempty" json:"shopid"`
	ListingID         int    `bson:"listingid,omitempty" json:"listingid"`
	TransactionID     int    `bson:"transactionid,omitempty" json:"transactionid"`
	BuyerUserID       int    `bson:"buyerid,omitempty" json:"buyerid"`
	Rating            int    `bson:"rating,omitempty" json:"rating"`
	Review            string `bson:"review,omitempty" json:"review"`
	Language          string `bson:"language,omitempty" json:"language"`
	ImageURLFullxfull string `bson:"imageurl,omitempty" json:"imageurl"`
	CreateTimestamp   int    `bson:"created,omitempty" json:"created"`
	UpdateTimestamp   int    `bson:"updated,omitempty" json:"updated"`
}

var ctx = context.Background()
var mongoUrl = os.Getenv("MONGOURL")
var opts = options.Aggregate().SetMaxTime(5 * time.Second)

var reviewsCollection = connect()

func connect() *mongo.Collection {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("bbpcontent")
	rc := database.Collection("reviews")
	return rc
}

func Random(w http.ResponseWriter, r *http.Request) {
	aggr := bson.D{{"$sample", bson.D{{"size", 9}}}}
	cursor, err := reviewsCollection.Aggregate(ctx, mongo.Pipeline{aggr}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var result []Review
	if err = cursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonResp)
}
