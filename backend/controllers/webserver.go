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
	subjectRouter.HandleFunc("/delete/{subject}", deleteSubject).Methods("POST", "OPTIONS")
	subjectRouter.HandleFunc("/edit/{subject}", updateSubject).Methods("POST", "OPTIONS")
	subjectRouter.HandleFunc("/random", getRandomSubject).Methods("GET", "OPTIONS")
	subjectRouter.HandleFunc("/{subject}", getSpecificSubject).Methods("GET", "OPTIONS")
	subjectRouter.Use(validateJWTMiddleware)
	studyLogRouter := router.PathPrefix("/api/study-log").Subrouter()
	studyLogRouter.HandleFunc("/start", addStudyStartLog)
	studyLogRouter.HandleFunc("/finish", addStudyFinishLog)
	studyLogRouter.HandleFunc("/all", getAllStudyLog)
	studyLogRouter.HandleFunc("/delete/{id}", deleteStudyLog)
	studyLogRouter.HandleFunc("/edit/{id}", updateStudyLog)
	studyLogRouter.HandleFunc("/subject", getSubjectStudyLog)
	studyLogRouter.HandleFunc("/{id}", getSpecificStudyLog)
	studyLogRouter.HandleFunc("/aggregation/daily", getDailyStudyInvestment)
	studyLogRouter.HandleFunc("/aggregation/weekly", getWeeklyStudyInvestment)
	studyLogRouter.HandleFunc("/aggregation/subject", getAggregationSubjectStudyTime)
	studyLogRouter.Use(validateJWTMiddleware)
	router.HandleFunc("/api/auth/register", registerUser).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/auth/user", GetTokenHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/auth/login", login).Methods("POST", "OPTIONS")
	fmt.Printf("Listen %s...\n", config.Config.ApiPort)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Config.ApiPort), router)
}
