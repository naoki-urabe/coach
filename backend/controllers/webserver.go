package controllers

import (
	"coach/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"math/rand"
	"coach/auth"
)

var addSubject = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	var subject models.Subject
	if err := json.Unmarshal(reqBody, &subject); err != nil {
		log.Fatal(err)
	}
	models.InsertSubject(&subject)
	responseBody, err := json.Marshal(subject)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})

var getAllSubject = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var subjects []models.Subject
	models.GetAllSubject(&subjects)
	responseBody, err := json.Marshal(subjects)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})

var getRandomSubject = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var subjects []models.Subject
	models.GetAllSubject(&subjects)
	n := len(subjects)
	i := rand.Intn(n-1)
	selectedSubject := subjects[i]
	responseBody, err := json.Marshal(selectedSubject)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Content-Type", "application/json")
}

func StartWebServer() error {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/api/subject", auth.JwtMiddleware.Handler(addSubject)).Methods("POST", "OPTIONS")
	router.Handle("/api/subject", auth.JwtMiddleware.Handler(getAllSubject)).Methods("GET")
	router.Handle("/api/subject/random", auth.JwtMiddleware.Handler(getRandomSubject)).Methods("GET")
	router.HandleFunc("/auth", auth.GetTokenHandler)
	fmt.Println("Listen 8080...")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
