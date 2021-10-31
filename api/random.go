package api

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goreviewapi/lib"
	"net/http"
)

func Random(w http.ResponseWriter, r *http.Request) {
	var ctx = context.Background()
	length := lib.GetLengthQuery(r)
	aggr := bson.D{{"$sample", bson.D{{"size", length}}}}
	cursor, err := lib.ReviewsCollection.Aggregate(ctx, mongo.Pipeline{aggr})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
