//simula ser un consumer en tiempo real

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Obj2 struct {
	Lista []Obj `json:"lista"`
}

type Obj struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Score string `json:"score"`
	Phase int    `json:"phase"`
	//Data  json.RawMessage `json:"data"`
}

func NuevoPartido(a string, b string, c string, d int) *Obj { //NuevaIndice
	return &Obj{a, b, c, d}
}

var collection = GetCollection("logs") //heroes

var ctx = context.Background()

func CreateV(carro Obj) error {

	var err error

	_, err = collection.InsertOne(ctx, carro)

	if err != nil {
		fmt.Println("@@@@@@@@@@@@@@@@@@@")
		fmt.Println(err)
		return err //si es error F
	}

	return nil
}

func Create(carro Obj) error {

	err := CreateV(carro)

	if err != nil {
		return err
	}

	return nil
}

func TestCreate(t *testing.T, carro Obj) {

	//carro := NuevoVehiculo("B618C", "Dodge", "Sedan", "Vyper", "Gris") //paso de Vehiculo a VehiculoMongoDB
	err := Create(carro)

	if err != nil {
		t.Error("Error en la prueba de persistencia de datos")
		fmt.Println(("Error en la prueba de persistencia de datos"))
		t.Fail()
	} else {
		t.Log("Todo salio a la perfeccion!!!!")
		fmt.Println("Todo salio a la perfeccion en create!!!!")

	}

}

var (
	/*usr      = "root"
	pwd      = "123456"
	host     = "database"
	port     = 27017*/
	database = "usactar"
	host     = "34.125.46.132" // ojo con ip dinamica TuT
	port     = 27017
)

func GetCollection(collection string) *mongo.Collection {

	//"mongodb://%s:%d", host, port						quickstart
	//"mongodb://%s:%s@%s:%d", usr, pwd, host, port		practica 1
	//uri := fmt.Sprintf("mongodb://%s:%d", host, port)

	uri := fmt.Sprintf("mongodb://basemongodb:J7ydEeQoOatZzLm9JuLTqrHhbxIb5Wh4PgrefmVMMyFnrV3vGC7H5i5Z169z25nUCHfvtlastx6zRdTGKBOJ7w==@basemongodb.mongo.cosmos.azure.com:10255/?ssl=true&retrywrites=false&replicaSet=globaldb&maxIdleTimeMS=120000&appName=@basemongodb@")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error()) //no puede seguir si pasa algo
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error()) //no puede seguir si pasa algo
	}

	fmt.Printf("coneccion correcta a coleccion:%s de la db:%s en mongoDB\n", collection, database)

	return client.Database(database).Collection(collection)
}

var t testing.T

//-------------------------------------

func main() {

	conf := kafka.ReaderConfig{
		Brokers:  []string{"my-cluster-kafka-bootstrap:9092"},
		Topic:    "my-topic",
		GroupID:  "g1",
		MaxBytes: 10000, //ojo
	}

	reader := kafka.NewReader(conf)

	fmt.Println("kafka listenner----------------\nfuck\n:(")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		fmt.Println("Msg: ", string(m.Value))
		//le envio a mongoDB y a Joel xd
		//-------------------------------- de []bytes a JSON--------------------
		//byt := []byte(`{"id":"someID","data":"hello world"}`)

		var obj Obj2
		if err := json.Unmarshal(m.Value, &obj); err != nil {
			panic(err)
		}

		fmt.Println(obj.Lista)
		//v1 := NuevoPartido("Guatemala", "Argentina", "0-3", 3)
		//joel

		for i := 0; i < len(obj.Lista); i++ {
			fmt.Printf(">>>escribi en mongo:%s vs %s\n", obj.Lista[i].Team1, obj.Lista[i].Team2)
			TestCreate(&t, obj.Lista[i]) //los registros por default
			//joel

			b, _ := json.Marshal(obj.Lista[i])
			enviarJoel(b)

		}

	}

}

func enviarJoel(buff []byte) {

	fmt.Println("recibilo Joel")
	responseBody := bytes.NewBuffer(buff)

	resp, err := http.Post("http://svc-grpc:8000/", "application/json", responseBody)

	if err != nil {
		log.Println("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	sb := string(body)
	log.Printf(sb)

}
