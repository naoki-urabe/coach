package controllers

import (
	"coach/models"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

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
	privateKeyPemStr := exportPrivateKeyAsPEMStr(privateKey)
	publicKeyPemStr := exportPublicKeyAsPEMStr(publicKey)
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
	subjectRouter.Use(validateJWTMiddleware)
	router.HandleFunc("/api/auth/register", registerUser).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/auth/user", GetTokenHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/auth/login", login).Methods("POST", "OPTIONS")
	fmt.Println("Listen 8080...")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
