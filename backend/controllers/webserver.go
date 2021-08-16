package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"coach/models"
)

func StartWebServer() error {
	models.ConnectDb()
	router := mux.NewRouter().StrictSlash(true)
	subjectRouter := router.PathPrefix("/api/subject").Subrouter()
	subjectRouter.HandleFunc("", addSubject).Methods("POST", "OPTIONS")
	subjectRouter.HandleFunc("", getAllSubject).Methods("GET")
	subjectRouter.HandleFunc("/random", getRandomSubject).Methods("GET", "OPTIONS")
	subjectRouter.Use(validateJWTMiddleware)
	router.HandleFunc("/api/auth/register", registerUser).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/auth/user", GetTokenHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/auth/login", login).Methods("POST", "OPTIONS")
	fmt.Println("Listen 8080...")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
