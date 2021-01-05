package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	db := InitDb()

	Route(r, db)

	// Look for port in env.json
	envFile, err := ioutil.ReadFile("env.json")

	if err != nil {
		panic("Error opening env.json!")
	}

	var jsonVal map[string]interface{}
	json.Unmarshal(envFile, &jsonVal)

	serverPort := jsonVal["server_port"]

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}).Handler(r)

	fmt.Println(fmt.Sprintf("Listening on http://localhost:%s!", serverPort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", serverPort), handler))
}
