package controllers

import (
	"coach/auth"
	"coach/models"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	mrand "math/rand"
	"net/http"
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
	i := mrand.Intn(n - 1)
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

func registerUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user models.User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		log.Fatal(err)
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey
	privateKeyPemStr := auth.ExportPrivateKeyAsPEMStr(privateKey)
	publicKeyPemStr := auth.ExportPublicKeyAsPEMStr(publicKey)
	user.PrivateKey = privateKeyPemStr
	user.PublicKey = publicKeyPemStr
	models.InsertUser(&user)
	responseBody, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
}

func StartWebServer() error {
	router := mux.NewRouter().StrictSlash(true)
	subjectRouter := router.PathPrefix("/api/subject").Subrouter()
	subjectRouter.HandleFunc("/", addSubject).Methods("POST", "OPTIONS")
	subjectRouter.HandleFunc("/", getAllSubject).Methods("GET")
	subjectRouter.HandleFunc("/random", getRandomSubject).Methods("GET", "OPTIONS")
	subjectRouter.Use(auth.ValidateJWTMiddleware)
	// router.HandleFunc("/auth", auth.GetTokenHandler)
	router.HandleFunc("/api/auth/register", registerUser).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/auth/user", auth.GetTokenHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/auth/login", auth.Login).Methods("POST", "OPTIONS")
	fmt.Println("Listen 8080...")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
