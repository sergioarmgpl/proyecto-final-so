package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var ctx = context.Background()

func url_root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "OK"}`))
}

type Response struct {
	N1 string `json:"n1"`
	N2 string `json:"n2"`
	Op string `json:"op"`
}

func write2Redis(n1 string, n2 string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "n1", n1, 0).Err()
	if err != nil {
		panic(err)
	}

	err2 := rdb.Set(ctx, "n2", n2, 0).Err()
	if err2 != nil {
		panic(err)
	}
}

func calculate() int {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	n1, err := rdb.Get(ctx, "n1").Result()
	if err != nil {
		panic(err)
	}

	n2, err2 := rdb.Get(ctx, "n2").Result()
	if err2 != nil {
		panic(err2)
	}

	return str2Int(n1) + str2Int(n2)
}

func str2Int(number string) int {
	if n, err := strconv.Atoi(number); err == nil {
		return n
	} else {
		return -1000
	}
}

func saveData(w http.ResponseWriter, r *http.Request) {

	n1Str := mux.Vars(r)["n1"]
	n2Str := mux.Vars(r)["n2"]
	op := "1"

	write2Redis(n1Str, n2Str)

	response := Response{n1Str, n2Str, op}
	fmt.Printf(response.N1, response.Op)

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func processData(w http.ResponseWriter, r *http.Request) {

	op := mux.Vars(r)["op"]
	result := calculate()
	result_a := strconv.Itoa(result)

	response := Response{result_a, "-1", op}

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
	router.HandleFunc("/save/n1/{n1}/n2/{n2}", saveData).Methods(http.MethodGet)
	router.HandleFunc("/process/op/{op}", processData).Methods(http.MethodGet)
	fmt.Println("Go ready!")
	log.Fatal(http.ListenAndServe(":5000", router))
}
