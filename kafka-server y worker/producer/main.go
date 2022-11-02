package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	/*
		El paquete gorilla/muximplementa un enrutador de solicitudes y un despachador para hacer coincidir las solicitudes entrantes con su respectivo controlador
	*/
	//"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	/*
		Un puerto Go (golang) del proyecto Ruby dotenv (que carga env vars desde un archivo .env)
	*/)

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

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
	}
}

func main() {
	//Mensaje de inicializacion
	fmt.Println("Servidor Kafka iniciado")
	fmt.Println("http://localhost:3000/")

	//Obligando que ese escriba la URL correcta
	router := mux.NewRouter().StrictSlash(true)

	//Definciendo la ruta, URL principal
	router.HandleFunc("/", inicio)

	//Definciendo la ruta, URL principal
	router.HandleFunc("/salud", salud2)
	router.HandleFunc("/input", createUser).Methods("POST")

	//Obtengo el puerto al cual se conecta desde el ".env"
	http_port := ":3000"

	//Levantamos el server
	//Por si existe un error
	if err := http.ListenAndServe(http_port, router); err != nil {
		fmt.Println(err)
	}
}

func inicio(w http.ResponseWriter, r *http.Request) { //creo q falta la conexion hacia el pod de kafka-topic o nose
	fmt.Fprintf(w, "sopes 1 grupo 5 \ncoordinador: Joel\n")

}

func salud2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aqui viene lo de kafka\n")
	//get_message(w)
}

func createUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var partidos Obj2
	json.NewDecoder((request.Body)).Decode(&partidos) //paso del objetoJson a obj de golang

	//muestro q ya los tengo como objeto-golang
	fmt.Println("------------recibio la api-go-ingress")
	for i := 0; i < len(partidos.Lista); i++ {
		fmt.Printf("Msg: %s vs %s (%s) fase:%d\n", partidos.Lista[i].Team1, partidos.Lista[i].Team2, partidos.Lista[i].Score, partidos.Lista[i].Phase)
	}

	//ahora q ya tengo los datos dentro de mi obj-go lo paso a []byte
	b, _ := json.Marshal(partidos)
	// muestra de q se convirtio a []byte
	s := string(b)
	fmt.Printf(">>>\nJSON -> []BYTE:\n%s", s)
	fmt.Println("--------------------\nsend message\n")

	//pilar para conectar a kafka

	send_message(b)
	//get_message(response)

	//retorno a postman thunder etc
	json.NewEncoder(response).Encode(partidos)
}

//no va a funcionar por la incopatibilidad de librerias xd -----------------------------------------------------------------------------------------

func send_message(datos []byte) { //modo ruso xd

	//configuracion inicial
	conn, err := kafka.DialLeader(context.Background(), "tcp", "my-cluster-kafka-bootstrap:9092", "my-topic", 0) //ojo cambiar al de gcp

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	//configuracion de tiempo
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	//configuracion del mensaje a ingresar a kafka
	conn.WriteMessages(kafka.Message{Value: datos})

}

/*
func get_message(w http.ResponseWriter) {

	//configuracion inicial
	conn, err := kafka.DialLeader(context.Background(), "tcp", "my-cluster-kafka-bootstrap:9092", "my-topic", 0) //ojo cambiar al de gcp

	if err != nil {
		fmt.Printf(err.Error())
	}

	//configuracion de tiempo
	//conn.SetWriteDeadline(time.Now().Add(time.Second * 8))
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	//configuracion de lectura de mensajes ingresados a kafka
	batch := conn.ReadBatch(0, 1e6) //1e3=10kb, 1e6 = 1mb max

	//       bytes := make([]byte, 1e6) //ojo con el tam 1e9 no lo soporta xd

	contador := 1

	for {
		bytes := make([]byte, 1e3) //ojo con el tam 1e9 no lo soporta xd
		_, err := batch.Read(bytes)
		if err == nil {
			fmt.Println("--------------------------------------------------------------------------------")
			fmt.Fprintf(w, "--------------------------------------------------------------------------------")
			fmt.Printf("\nMsg%d:%s\n", contador, string(bytes))
			fmt.Fprintf(w, "\nMsg%d:%s\n", contador, string(bytes))
			contador++
			//fmt.Println("--------------------------------------------------------------------------------")
			//time.Sleep(1 * time.Second)

		} else { //error
			break
		}
		//fmt.Println("Msg:",string(bytes))

	}

	//validaciones del repo oficial

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

*/
