package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	//"io/ioutil"
	"log"
	"net/http"

	"os"
	"time"

	pb "client/proto-grpc"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type Informacion struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Score string `json:"score"`
	Phase int    `json:"phase"`
}

func conectar_server(wri http.ResponseWriter, req *http.Request) {
	//Declaraciones
	var info Informacion

	//Obtener direccion del server
	host := os.Getenv("HOST_GRPC")
	//host = "localhost"
	// AGREGAR CORS
	SetCors(wri)

	conn, err := grpc.Dial(host+":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		json.NewEncoder(wri).Encode("No se puede conectar con el server de grpc")
		log.Fatalf("No se pudo conectar con el server :c (%v)", err)
	}
	//Cerrar el servidor
	defer conn.Close()

	//Iniciar objeto
	cl := pb.NewGetInfoClient(conn)
	//Get info enviada para guardar
	json.NewDecoder(req.Body).Decode(&info)
	//Crear contexto
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	//Mandar info al servidor
	ret, err := cl.ReturnInfo(ctx, &pb.RequestInfo{Team1: info.Team1, Team2: info.Team2,
		Score: info.Score, Phase: strconv.Itoa(info.Phase)})
	if err != nil {
		json.NewEncoder(wri).Encode("Error, no  se puede retornar la información.")
		log.Fatalf("No se puede retornar la información :c (%v)", err)
	}
	//Retornar respuesta
	log.Printf("Respuesta del server: %s\n", ret.GetInfo())
	json.NewEncoder(wri).Encode(ret.GetInfo())
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", conectar_server).Methods("POST")
	fmt.Println("Cliente se levanto en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func SetCors(wri http.ResponseWriter) {
	(wri).Header().Set("Access-Control-Allow-Origin", "*")
	(wri).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(wri).Header().Set("Content-Type", "application/json")
}
