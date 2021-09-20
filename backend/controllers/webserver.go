package controllers

import (
	"coach/config"
	"coach/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func StartWebServer() error {
	models.ConnectDb()
	router := mux.NewRouter().StrictSlash(true)
	subjectRouter := router.PathPrefix("/api/subject").Subrouter()
	subjectRouter.HandleFunc("", addSubject).Methods("POST", "OPTIONS")
	subjectRouter.HandleFunc("", getAllSubject).Methods("GET")
	subjectRouter.HandleFunc("/random", getRandomSubject).Methods("GET", "OPTIONS")
	subjectRouter.Use(validateJWTMiddleware)
	studyLogRouter := router.PathPrefix("/api/study-log").Subrouter()
	studyLogRouter.HandleFunc("/start", addStudyStartLog)
	studyLogRouter.HandleFunc("/finish", addStudyFinishLog)
	studyLogRouter.HandleFunc("/all", getAllStudyLog)
	studyLogRouter.HandleFunc("/aggregation/daily", getDailyStudyInvestment)
	studyLogRouter.HandleFunc("/aggregation/weekly", getWeeklyStudyInvestment)
	studyLogRouter.Use(validateJWTMiddleware)
	router.HandleFunc("/api/auth/register", registerUser).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/auth/user", GetTokenHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/auth/login", login).Methods("POST", "OPTIONS")
	fmt.Printf("Listen %s...", config.Config.ApiPort)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Config.ApiPort), router)
}
