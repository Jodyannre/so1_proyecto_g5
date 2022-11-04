package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "server/proto-grpc"

	//"time"
	"github.com/go-redis/redis/v9"
	"google.golang.org/grpc"
)

var (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

func almacenar_datos(t1 string, t2 string, field string, phase string) {
	//Guardar informacion en redis
	fmt.Println("Esta es la info que se guardaria en la BD")
	var key = fmt.Sprintf("%s:%s:%s", t1, t2, phase)
	var equipos = fmt.Sprintf("%s:%s", t1, t2)
	var fase = fmt.Sprintf("Fase:%s", phase)
	var keyTotal = fmt.Sprintf("%s:%s:%s:total", t1, t2, phase)
	fmt.Printf(key + "\n")

	//Crear la conexion
	//Crear la conexion
	var ctx = context.Background()
	const ip = "baseredisdb2.redis.cache.windows.net:6379"
	rdb := redis.NewClient(&redis.Options{
		Addr:     ip,
		Password: "Z4OB967BvclEP8d11anmobzkHx92EXSEWAzCaC0nUiY=",
		DB:       0,
	})
	//Agregar o incrementar  valor
	val := rdb.HIncrBy(ctx, key, field, 1)
	total := rdb.IncrBy(ctx, keyTotal, 1)
	conjunto := rdb.SAdd(ctx, fase, equipos)
	miemros := rdb.SMembers(ctx, "Fase:3")
	fmt.Println(miemros)
	//ErrorHandler(err)
	fmt.Println("El valor guardado  es: ", val)
	fmt.Println("El total es: ", total)
	fmt.Println("El estado del conjunto es: ", conjunto)
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestInfo) (*pb.ReplyInfo, error) {
	almacenar_datos(in.GetTeam1(), in.GetTeam2(), in.GetScore(), in.GetPhase())
	resultado := fmt.Sprintf("%s vs %s : %s phase: %s", in.GetTeam1(), in.GetTeam2(), in.GetScore(), in.GetPhase())
	return &pb.ReplyInfo{Info: ">> Datos guardados: " + resultado}, nil
}

func main() {
	escucha, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Fallo al levantar el servidor: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})
	if err := s.Serve(escucha); err != nil {
		log.Fatalf("Fallo al levantar el servidor: %v", err)
	}
}

func ErrorHandler(err any) {
	if err != nil {
		panic(err)
	}
}

/*
func redis(){
	var ctx = context.Background()
	ip := "172.17.0.2:6379"
	rdb := redis.NewClient(&redis.Options{
		Addr: ip,
		Password:"",
		DB:0,
	})
	err := rdb.Set(ctx, "key", "value",0).Err()
	if err != nil{
		panic(err)
	}
	val,err := rdb.Get(ctx, "key").Result()
	if err != nil{
		panic(err)
	}
	fmt.Println("key",val)
}
*/
