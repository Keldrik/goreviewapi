package api

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goreviewapi/lib"
	"net/http"
)

func Listing(w http.ResponseWriter, r *http.Request) {
	var ctx = context.Background()
	length := lib.GetLengthQuery(r)
	listing := lib.GetListingQuery(r)
	if listing == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	aggr1 := bson.D{{"$match", bson.D{{"listingid", listing}}}}
	aggr2 := bson.D{{"$sample", bson.D{{"size", length}}}}
	cursor, err := lib.ReviewsCollection.Aggregate(ctx, mongo.Pipeline{aggr1, aggr2})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if cursor.RemainingBatchLength() == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var result []lib.Review
	if err = cursor.All(ctx, &result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonResp, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
