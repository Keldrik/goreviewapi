package api

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goreviewapi/lib"
	"log"
	"net/http"
	"strconv"
)

func Listing(w http.ResponseWriter, r *http.Request) {
	var ctx = context.Background()
	listingQuery, ok := r.URL.Query()["listing"]
	if !ok || len(listingQuery[0]) < 1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	listing, _ := strconv.Atoi(listingQuery[0])
	aggr1 := bson.D{{"$match", bson.D{{"listingid", listing}}}}
	aggr2 := bson.D{{"$sample", bson.D{{"size", 6}}}}
	cursor, err := lib.ReviewsCollection.Aggregate(ctx, mongo.Pipeline{aggr1, aggr2})
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
