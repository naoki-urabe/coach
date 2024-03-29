package models

import (
	"log"
	mrand "math/rand"
)

type Subject struct {
	SubjectCode string `db:"subject_code" json:"subject_code"`
	SubjectName string `db:"subject_name" json:"subject_name"`
}

type Kuji struct {
	SubjectCode string `db:"subject_code" json:"subject_code"`
	Selected    bool   `db:"selected" json:"selected"`
	Target      bool   `db:"target" json:"target"`
}

var getAllSubjectQuery = `
SELECT * FROM subjects;
`

var getKujiAllQuery = `
SELECT * FROM kuji;
`

var getSpecificSubjectQuery = `
SELECT * FROM subjects WHERE subject_code = ?;
`

var getSubjectKujiQuery = `
SELECT * FROM kuji WHERE selected = 0 AND target = 1;
`

var updateKujiQuery = `
UPDATE kuji SET selected = 1 WHERE subject_code = ?;
`

var updateKujiReset = `
UPDATE kuji SET selected = 0;
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

func GetSubjectKuji(subject *Subject) {
	var kuji []Kuji
	Db.Select(&kuji, getSubjectKujiQuery)
	n := len(kuji)
	if n == 0 {
		_, err := Db.Queryx(updateKujiReset)
		Db.Select(&kuji, getSubjectKujiQuery)
		n = len(kuji)
		if err != nil {
			log.Println(err)
		}
	}
	i := mrand.Intn(n)
	_, err := Db.Queryx(updateKujiQuery, kuji[i].SubjectCode)
	if err != nil {
		log.Println(err)
	}
	Db.Get(subject, getSpecificSubjectQuery, kuji[i].SubjectCode)
}

func GetKujiAll(kuji *[]Kuji) {
	Db.Select(kuji, getKujiAllQuery)
}
