package models

type Subject struct {
	SubjectCode string `db:"subject_code" json:"subject_code"`
	Subject     string `db:"subject" json:"subject"`
}

var getAllSubjectQuery = `
SELECT * FROM subjects;
`

var insertSubjectQuery = `
INSERT INTO subjects VALUES (?, ?);`

func GetAllSubject(subject *[]Subject) {
	Db.Select(subject, getAllSubjectQuery)
}

func InsertSubject(subject *Subject) {
	Db.MustExec(insertSubjectQuery)
}
