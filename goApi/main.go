package main

import (
	"context"
	"fmt"
	mongodb "goApi/mongodb"
	redisdb "goApi/redisdb"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//Carger el arvhio env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	PORT_APP := os.Getenv("API_PORT")
	//Crear conexión de mongo
	//uri := fmt.Sprintf("mongodb://%s:%s", hostMongo, portMongo)
	uri := os.Getenv("MONGO_HOST")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientMongo, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	//Crear conexión de redis
	hostRedis := os.Getenv("REDIS_HOST")
	portRedis := os.Getenv("REDIS_PORT")
	passRedis := os.Getenv("REDIS_PASS")

	clientRedis := redis.NewClient(&redis.Options{
		Addr:     hostRedis + ":" + portRedis,
		Password: passRedis,
		DB:       0,
	})
	var ctx2 = context.Background()

	router := mux.NewRouter()
	router.HandleFunc("/getLogsMongo", mongodb.GetLogsMongo(clientMongo)).Methods("GET")
	router.HandleFunc("/logsMongo", mongodb.WriteLogMongo(clientMongo)).Methods("POST")
	router.HandleFunc("/getTotalMongo", mongodb.GetTotalMongo(clientMongo)).Methods("GET")
	router.HandleFunc("/saveMatch", redisdb.SaveMatch(clientRedis, ctx2)).Methods("POST")
	router.HandleFunc("/getDataFases", redisdb.GetAllMatches(clientRedis, ctx2)).Methods("GET")
	router.HandleFunc("/getDataPartidos", redisdb.GetCounters(clientRedis, ctx2)).Methods("GET")

	fmt.Printf("App running at port: %s", PORT_APP)
	http.ListenAndServe(":"+PORT_APP, router)
}
