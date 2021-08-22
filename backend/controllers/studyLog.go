package controllers

import (
	"coach/models"
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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
	responseBody, err := json.Marshal(finishLog)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})
