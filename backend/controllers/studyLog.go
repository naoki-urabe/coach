package controllers

import (
	"coach/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	User string
}

var addStudyStartLog = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	var startLog models.StartLog
	if err := json.Unmarshal(reqBody, &startLog); err != nil {
		log.Fatal(err)
	}
	if startLog.SubjectCode == "" {
		w.WriteHeader(406)
		return
	}
	id := models.AddStudyStartLog(&startLog)
	startLog.Id = id
	responseBody, err := json.Marshal(startLog)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})

var addStudyFinishLog = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	var finishLog models.FinishLog
	if err := json.Unmarshal(reqBody, &finishLog); err != nil {
		log.Fatal(err)
	}
	models.AddStudyFinishLog(&finishLog)
	var latestStudyLog models.StudyLog
	models.GetLatestStudyLog(&latestStudyLog, finishLog.Id)
	responseBody, err := json.Marshal(latestStudyLog)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})

var getAllStudyLog = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	var user User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		panic(err)
	}
	var studyLogs []models.StudyLog
	models.GetAllStudyLog(user.User, &studyLogs)
	responseBody, err := json.Marshal(studyLogs)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})

var deleteStudyLog = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	vars := mux.Vars(r)
	deleteId, _ := strconv.Atoi(vars["id"])
	models.DeleteStudyLog(deleteId)
	fmt.Fprintln(w, vars["id"])
})
var getSubjectStudyLog = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	type Req struct {
		User        string
		SubjectCode string
	}
	var req Req
	if err := json.Unmarshal(reqBody, &req); err != nil {
		panic(err)
	}
	var studyLogs []models.StudyLog
	models.GetSubjectStudyLog(req.User, req.SubjectCode, &studyLogs)
	responseBody, err := json.Marshal(studyLogs)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})
var getDailyStudyInvestment = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	var user User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		panic(err)
	}
	var periodDiff []models.PeriodDiff
	models.GetDailyStudyInvestment(user.User, &periodDiff)
	responseBody, err := json.Marshal(periodDiff)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})

var getWeeklyStudyInvestment = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	var user User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		panic(err)
	}
	var periodDiff []models.PeriodDiff
	models.GetWeeklyStudyInvestment(user.User, &periodDiff)
	responseBody, err := json.Marshal(periodDiff)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})
