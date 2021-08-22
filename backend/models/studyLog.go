package models

import (
	_ "fmt"
	"time"
)

var addStudyStartLogQuery = `
INSERT INTO study_logs (subject_code,study_start_time) VALUES(?,?);
`

var addStudyFinishLogQuery = `
UPDATE study_logs SET comment = ?,study_finish_time = ? WHERE id = ?;
`

/*type StudyLog struct {
	Id int `db:"id" json:"id"`
	SubjectCode string `db:"subject_code" json:"subject_code"`
	Comment string `db:"comment" json:`
	StudyStartTime time.Time `db:"study_start_time" json:"study_start_time"`
	StudyFinishTime time.Time `db:"study_finish_time" json:"study_finish_time"`
}*/

type StartLog struct {
	Id             int       `db:"id" json:"id"`
	SubjectCode    string    `db:"subject_code" json:"subject_code"`
	StudyStartTime time.Time `db:"study_start_time" json:"study_start_time"`
}

type FinishLog struct {
	Id              int       `db:"id" json:"id"`
	Comment         string    `db:"comment" json:"comment"`
	StudyFinishTime time.Time `db:"study_finish_time" json:"study_finish_time"`
}

func AddStudyStartLog(startLog *StartLog) int {
	registerLog := Db.MustExec(addStudyStartLogQuery, startLog.SubjectCode, startLog.StudyStartTime)
	id, _ := registerLog.LastInsertId()
	return int(id)
}

func AddStudyFinishLog(finishLog *FinishLog) {
	Db.MustExec(addStudyFinishLogQuery, finishLog.Comment, finishLog.StudyFinishTime, finishLog.Id)
}
