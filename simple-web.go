package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Member struct {
	Name  string
	Age   int
	Alive bool
}

func info(w http.ResponseWriter, r *http.Request) {
	mem := Member{"Heo", 29, true}
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}

	//w.Header().Set("Content-Type", "application/json")
	//w.Write(jsonBytes)
	jsonString := string(jsonBytes)
	fmt.Fprintf(w, jsonString)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func sum(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	a, _ := strconv.Atoi(v["a"])
	b, _ := strconv.Atoi(v["b"])
	fmt.Fprintf(w, "%d", a+b)
}

func main() {
	h := mux.NewRouter()
	h.HandleFunc("/hello", hello).Methods("GET")
	h.HandleFunc("/sum/{a}/{b}", sum).Methods("GET")
	h.HandleFunc("/info", info).Methods("GET")
	http.Handle("/", h)
	http.ListenAndServe(":3000", nil)
}
