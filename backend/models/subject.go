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

var getSpecificSubjectQuery = `
SELECT * FROM subjects WHERE subject_code = ?;
`

var insertSubjectQuery = `
INSERT INTO subjects VALUES(?, ?);`

var updateSubjectQuery = `
UPDATE subjects SET subject_code = ?, subject_name = ? WHERE subject_code = ?;`

var deleteSubjectQuery = `
DELETE FROM subjects WHERE subject_code = ?;
`

func GetAllSubject(subject *[]Subject) {
	Db.Select(subject, getAllSubjectQuery)
}

func GetSpecificSubject(subject *Subject, subjectCode string) {
	Db.Get(subject, getSpecificSubjectQuery, subjectCode)
}

func InsertSubject(subject *Subject) bool {
	_, err := Db.Queryx(insertSubjectQuery, subject.SubjectCode, subject.SubjectName)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func UpdateSubject(subject *Subject, subjectCode string) {
	_, err := Db.Queryx(updateSubjectQuery, subject.SubjectCode, subject.SubjectName, subjectCode)
	if err != nil {
		log.Println(err)
	}
}

func DeleteSubject(subject string) {
	_, err := Db.Queryx(deleteSubjectQuery, subject)
	if err != nil {
		log.Println(err)
	}
}
