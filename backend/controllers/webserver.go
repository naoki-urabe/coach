package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

func StartWebServer() error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", hello).Methods("GET")
	fmt.Println("Listen 8080...")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}