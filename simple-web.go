package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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
	http.Handle("/", h)
	http.ListenAndServe(":3000", nil)
}
