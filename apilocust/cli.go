package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	Ejecutar()
}

func Ejecutar() {
	consola := color.New(color.FgBlue).Add(color.Bold)
	errorconsola := color.New(color.FgRed).Add(color.Bold)
	exitoconsola := color.New(color.FgGreen).Add(color.Bold)
	consola.Println("********************************************")
	consola.Println("**** PROYECTO 2 - SISTEMAS OPERATIVOS 2 ****")
	consola.Println("                                            ")
	consola.Println("****** CLI PARA EJECUTAR LOCUST EN GO ******")
	consola.Println("********************************************")
	comando := color.New(color.BgGreen).Add(color.Bold)
	comando.Println("Ingrese el comando: run -f data.json --concurrence <no.Concurrencia> -n <no.Peticiones> --timeout <no.segundos>")
	//ingresar comando
	var comandoIngresado string
	lectura := bufio.NewReader(os.Stdin)
	comandoIngresado, _ = lectura.ReadString('\n')
	comandoIngresado = strings.Replace(comandoIngresado, "\n", "", -1) //elimino salto
	comandoIngresado = strings.Replace(comandoIngresado, "\r", "", -1) //elimino salto
	//ahora se separa el comando ingresado
	comandoSeparado := separarComando(comandoIngresado)
	//se ejecuta el comando
	fmt.Println(comandoSeparado)
	//validar tamaño del comando
	if len(comandoSeparado) < 9 {
		errorconsola.Println("Comando ingresado no es valido")
	} else {
		//validar comando
		if comandoSeparado[0] == "run" && comandoSeparado[1] == "-f" && comandoSeparado[3] == "--concurrence" && comandoSeparado[5] == "-n" && comandoSeparado[7] == "--timeout" {
			//validar que el archivo exista
			if _, err := os.Stat(comandoSeparado[2]); os.IsNotExist(err) {
				errorconsola.Println("El archivo no existe")
			} else {
				//validar que el archivo sea json
				if strings.HasSuffix(comandoSeparado[2], ".json") {
					//validar que el numero de concurrencia sea un numero
					if _, err := strconv.Atoi(comandoSeparado[4]); err == nil {
						//validar que el numero de peticiones sea un numero
						if _, err := strconv.Atoi(comandoSeparado[6]); err == nil {
							//validar que el timeout sea un numero
							if _, err := strconv.Atoi(comandoSeparado[8]); err == nil {
								//ejecutar el comando
								exitoconsola.Println("Ejecutando comando")
								//se ejecuta el comando
								ejecutarComando(comandoSeparado[2], comandoSeparado[4], comandoSeparado[6], comandoSeparado[8])
							} else {
								errorconsola.Println("El timeout no es un numero")
							}
						} else {
							errorconsola.Println("El numero de peticiones no es un numero")
						}
					} else {
						errorconsola.Println("El numero de concurrencia no es un numero")
					}
				} else {
					errorconsola.Println("El archivo no es un archivo json")
				}
			}
		} else {
			errorconsola.Println("Comando ingresado no es valido")
		}
	}

	//comandoSeparado := separarComando(entrada)
	//ejecutar comando
	//cmd := exec.Command("locust", "-f", "traffic.py", "--host=http://34.75.240.99.nip.io", "--headless", "-u", "10", "-r", "1")
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//cmd.Run()

}

func ejecutarComando(archivo string, concurrencia string, peticiones string, timeout string) {
	//se ejecuta el comando
	cmd := exec.Command("locust", "-f", "traffic.py", "--host=http://36.196.23.228.nip.io", "--headless", "-u", concurrencia, "-r", "1", "-t", timeout, "--my-argument", peticiones)
	//cmd := exec.Command("python", "main.py", "100")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func separarComando(comando string) []string {
	comandoSeparado := strings.Split(comando, " ")
	return comandoSeparado

}

// ejemplo comando
// run -f data.json --concurrence 10 -n 20 --timeout 30000
