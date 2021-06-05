package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func url_root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "OK"}`))
}

func processData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`"output": { "health":"OK"}`))
	n1 := mux.Vars(r)["n1"]
	fmt.Printf(n1)
	var1 := mux.Vars(r)["n2"]
	fmt.Printf(var1)
	op := mux.Vars(r)["op"]
	fmt.Printf(op)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", url_root).Methods(http.MethodGet)
	router.HandleFunc("/process/n1/{n1}/n2/{n2}/op/{op}", processData).Methods(http.MethodGet)
	fmt.Println("Go ready!")
	log.Fatal(http.ListenAndServe(":5000", router))
}
