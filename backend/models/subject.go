package models

import (
	"log"
)

type Subject struct {
	SubjectCode string `db:"subject_code" json:"subject_code"`
	SubjectName string `db:"subject_name" json:"subject_name"`
}

var getAllSubjectQuery = `
SELECT * FROM subjects;
`

var insertSubjectQuery = `
INSERT INTO subjects VALUES(?, ?);`

func GetAllSubject(subject *[]Subject) {
	Db.Select(subject, getAllSubjectQuery)
}

func InsertSubject(subject *Subject) bool {
	_, err := Db.Queryx(insertSubjectQuery, subject.SubjectCode, subject.SubjectName)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
