package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		return
	}
	fmt.Println("hello")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Content-Type", "application/json")
}

func StartWebServer() error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", hello).Methods("GET")
	fmt.Println("Listen 8080...")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}