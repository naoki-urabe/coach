package controllers

import (
	"coach/models"
	"encoding/json"
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
	i := mrand.Intn(n)
	selectedSubject := subjects[i]
	responseBody, err := json.Marshal(selectedSubject)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})
