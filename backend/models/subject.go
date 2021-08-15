package models

type Subject struct {
	SubjectCode string `db:"subject_code" json:"subject_code"`
	SubjectName     string `db:"subject_name" json:"subject_name"`
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
