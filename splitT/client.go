package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var cambia = true

func conectar_server(wri http.ResponseWriter, req *http.Request) {

	fmt.Println("Ingreso Solicitud>")
	// AGREGAR CORS
	wri.Header().Set("Access-Control-Allow-Origin", "*")
	wri.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	wri.Header().Set("Content-Type", "application/json")

	if req.Method == "GET" {
		wri.WriteHeader(http.StatusOK)
		wri.Write([]byte("{\"mensaje\": \"get\"}"))
		return
	}

	datos, _ := ioutil.ReadAll(req.Body)

	//inicia parte para consumir la api------------------------------------
	//https://httpbin.org/post2
	var res string
	if cambia {

		resp, err := http.Post(os.Getenv("RUTA1"), "application/json", bytes.NewBuffer(datos))

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		res = string(bodyBytes)

		fmt.Println("INFORMACION1>" + string(datos))
		cambia = false

	} else {
		resp, err := http.Post(os.Getenv("RUTA2"), "application/json", bytes.NewBuffer(datos))
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		res = string(bodyBytes)

		fmt.Println("INFORMACION2>" + string(datos))
		cambia = true
	}

	//fin para consumir api----------------------------------------------

	fmt.Println(res)
	json.NewEncoder(wri).Encode(res)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", conectar_server)
	fmt.Println("Cliente se levanto en el puerto " + os.Getenv("PUERTO"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PUERTO"), router))
}
