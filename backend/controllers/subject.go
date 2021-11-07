package controllers

import (
	"coach/models"
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
	if subject.SubjectName == "" || subject.SubjectCode == "" {
		w.WriteHeader(406)
		log.Printf("Error Register Subject:%s\n", subject.SubjectName)
		return
	}
	isSuccess := models.InsertSubject(&subject)
	responseBody, err := json.Marshal(subject)
	if err != nil {
		log.Fatal(err)
	}
	if isSuccess {
		log.Printf("Success Register Subject:%s\n", subject.SubjectName)
		w.Write(responseBody)
	} else {
		log.Printf("Error Register Subject:%s\n", subject.SubjectName)
		w.WriteHeader(409)
	}
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

var getSpecificSubject = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	var subject models.Subject
	vars := mux.Vars(r)
	specificSubject := vars["subject"]
	models.GetSpecificSubject(&subject, specificSubject)
	responseBody, err := json.Marshal(subject)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})
var updateSubject = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	vars := mux.Vars(r)
	updateSubject := vars["subject"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var subject models.Subject
	if err := json.Unmarshal(reqBody, &subject); err != nil {
		log.Fatal(err)
	}
	models.UpdateSubject(&subject, updateSubject)
	fmt.Fprintln(w, updateSubject)
})

var deleteSubject = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	vars := mux.Vars(r)
	deleteSubject := vars["subject"]
	models.DeleteSubject(deleteSubject)
	fmt.Fprintln(w, deleteSubject)
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
