package main

import (
	"encoding/json"
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

type Response struct {
	N1 string
	N2 string
	Op string
}

func processData(w http.ResponseWriter, r *http.Request) {

	n1 := mux.Vars(r)["n1"]
	fmt.Printf(n1)
	n2 := mux.Vars(r)["n2"]
	fmt.Printf(n2)
	op := mux.Vars(r)["op"]
	fmt.Printf(op)

	response := Response{n1, n2, op}
	fmt.Printf(response.N1, response.Op)

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", url_root).Methods(http.MethodGet)
	router.HandleFunc("/process/n1/{n1}/n2/{n2}/op/{op}", processData).Methods(http.MethodGet)
	fmt.Println("Go ready!")
	log.Fatal(http.ListenAndServe(":5000", router))
}
