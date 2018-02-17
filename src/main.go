package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/health", health).Methods("GET")

	fmt.Println("Running server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func health(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{"status": "OK"}
	json.NewEncoder(w).Encode(m)
}
