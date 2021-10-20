package api

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goreviewapi/lib"
	"log"
	"net/http"
)

func Random(w http.ResponseWriter, r *http.Request) {
	var ctx = context.Background()
	aggr := bson.D{{"$sample", bson.D{{"size", 9}}}}
	cursor, err := lib.ReviewsCollection.Aggregate(ctx, mongo.Pipeline{aggr})
	if err != nil {
		log.Fatal(err)
	}
	var result []lib.Review
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
