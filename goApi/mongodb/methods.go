package mongodb

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetLogsMongo(client *mongo.Client) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		SetCors(res)
		var logs LogsMongo
		collection := client.Database("usactar").Collection("logs")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		opts := options.Find().SetSort(bson.D{{"_id", -1}})
		opts.SetLimit(10)

		cursor, err := collection.Find(ctx, bson.M{}, opts)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(`{"mensaje": "` + err.Error() + `"}`))
			return
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var nLog LogMongo
			cursor.Decode(&nLog)
			logs = append(logs, nLog)
		}
		if err := cursor.Err(); err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(`{"mensaje": "` + err.Error() + `"}`))
			return
		}
		json.NewEncoder(res).Encode(logs)
	}
}

func GetTotalMongo(client *mongo.Client) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		SetCors(res)
		var total TotalMongo
		collection := client.Database("usactar").Collection("logs")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		result, err := collection.CountDocuments(ctx, bson.M{})
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(`{"mensaje": "` + err.Error() + `"}`))
			return
		}
		total.Total = result
		json.NewEncoder(res).Encode(total)

	}
}

func WriteLogMongo(client *mongo.Client) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		SetCors(res)
		var log LogMongo
		json.NewDecoder(req.Body).Decode(&log)
		collection := client.Database("usactar").Collection("logs")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		result, _ := collection.InsertOne(ctx, log)
		json.NewEncoder(res).Encode(result)
	}
}

func SetCors(wri http.ResponseWriter) {
	(wri).Header().Set("Access-Control-Allow-Origin", "*")
	(wri).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(wri).Header().Set("Content-Type", "application/json")
}
