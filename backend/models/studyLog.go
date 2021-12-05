package models

import (
	_ "fmt"
	"log"
	"time"
)

var addStudyStartLogQuery = `
INSERT INTO study_logs (subject_code,study_start_time,user) VALUES(?,?,?);
`

var addStudyFinishLogQuery = `
UPDATE study_logs SET content = ?,comment = ?,study_finish_time = ? WHERE id = ?;
`

var getAllStudyLogQuery = `
SELECT * FROM study_logs WHERE user = ?;
`

var getSubjectStudyLogQuery = `
SELECT * FROM study_logs WHERE user = ? AND subject_code = ?;
`

var getLatestStudyLogQuery = `
SELECT * FROM study_logs WHERE id = ?;`

var getDailyStudyInvestmentQuery = `
SELECT period,SUM(diff) AS diff
    FROM 
	    (SELECT CAST(study_start_time AS DATE) AS period,
		CAST((TIME_TO_SEC(TIMEDIFF(study_finish_time,study_start_time)) / 60) AS SIGNED) AS diff 
        FROM study_logs
		WHERE user = ?) 
	AS t
GROUP BY period;
`

var getWeeklyStudyInvestmentQuery = `
SELECT period, SUM(diff) AS diff
    FROM 
        (SELECT 
            CAST(SUBDATE(study_start_time, WEEKDAY(study_start_time)) as DATE) as period,
            CAST((TIME_TO_SEC(TIMEDIFF(study_finish_time,study_start_time)) / 60) AS SIGNED) AS diff 
            From study_logs
		WHERE user = ?)
        AS t
GROUP BY period;
`

var deleteStudyLogQuery = `
DELETE FROM study_logs WHERE id = ?;
`

var getSpecificStudyLogQuery = `
SELECT * FROM study_logs WHERE id = ?;
`

var updateStudyLogQuery = `
UPDATE study_logs SET subject_code = ?, content = ?, comment = ?,study_start_time = ?, study_finish_time = ? WHERE id = ?
`

var getAggregationSubjectStudyTimeQuery = `
SELECT SUM(CAST((TIME_TO_SEC(TIMEDIFF(study_finish_time,study_start_time)) / 60) AS SIGNED) ) AS aggregation_subject_study_time,subject_code FROM study_logs GROUP BY subject_code;
`

type StudyLog struct {
	Id              int        `db:"id" json:"id"`
	SubjectCode     string     `db:"subject_code" json:"subject_code"`
	Content         *string    `db:"content" json:"content"`
	Comment         *string    `db:"comment" json:"comment"`
	StudyStartTime  time.Time  `db:"study_start_time" json:"study_start_time"`
	StudyFinishTime *time.Time `db:"study_finish_time" json:"study_finish_time"`
	User            string     `db:"user" json:"user"`
}

type StartLog struct {
	Id             int       `db:"id" json:"id"`
	SubjectCode    string    `db:"subject_code" json:"subject_code"`
	StudyStartTime time.Time `db:"study_start_time" json:"study_start_time"`
	User           string    `db:"user" json:"user"`
}

type FinishLog struct {
	Id              int       `db:"id" json:"id"`
	Content         string    `db:"content" json:"content"`
	Comment         string    `db:"comment" json:"comment"`
	StudyFinishTime time.Time `db:"study_finish_time" json:"study_finish_time"`
}

type PeriodDiff struct {
	Period string `db:"period" json:"period"`
	Diff   string `db:"diff" json:"diff"`
}

type AggregationSubjectStudyTime struct {
	SubjectCode                 string `db:"subject_code" json:"subject_code"`
	AggregationSubjectStudyTime string `db:"aggregation_subject_study_time" json: "aggregation_subject_study_time"`
}

func AddStudyStartLog(startLog *StartLog) int {
	registerLog := Db.MustExec(addStudyStartLogQuery, startLog.SubjectCode, startLog.StudyStartTime, startLog.User)
	id, _ := registerLog.LastInsertId()
	return int(id)
}

func AddStudyFinishLog(finishLog *FinishLog) {
	Db.MustExec(addStudyFinishLogQuery, finishLog.Content, finishLog.Comment, finishLog.StudyFinishTime, finishLog.Id)
}

func UpdateStudyLog(studyLog *StudyLog, updateId int) {
	_, err := Db.Queryx(updateStudyLogQuery, studyLog.SubjectCode, studyLog.Content, studyLog.Comment, studyLog.StudyStartTime, studyLog.StudyFinishTime, updateId)
	if err != nil {
		log.Println(err)
	}
}

func DeleteStudyLog(deleteId int) {
	_, err := Db.Queryx(deleteStudyLogQuery, deleteId)
	if err != nil {
		log.Printf("DeleteStudyLog:%s\n", err)
	}
}

func GetAllStudyLog(user string, studyLog *[]StudyLog) {
	Db.Select(studyLog, getAllStudyLogQuery, user)
}

func GetSubjectStudyLog(user string, subjectCode string, studyLog *[]StudyLog) {
	err := Db.Select(studyLog, getSubjectStudyLogQuery, user, subjectCode)
	if err != nil {
		log.Printf("GetSubjectStudyLog:%s\n", err)
	}
}

func GetSpecificStudyLog(studyLog *StudyLog, id int) {
	err := Db.Get(studyLog, getSpecificStudyLogQuery, id)
	if err != nil {
		log.Printf("GetSpecificStudyLog:%s\n", err)
	}
}

func GetLatestStudyLog(studyLog *StudyLog, id int) {
	Db.Get(studyLog, getLatestStudyLogQuery, id)
}

func GetDailyStudyInvestment(user string, periodDiff *[]PeriodDiff) {
	Db.Select(periodDiff, getDailyStudyInvestmentQuery, user)
}

func GetWeeklyStudyInvestment(user string, periodDiff *[]PeriodDiff) {
	Db.Select(periodDiff, getWeeklyStudyInvestmentQuery, user)
}
func GetAggregationSubjectStudyTime(aggSubjectStudyTime *[]AggregationSubjectStudyTime) {
	err := Db.Select(aggSubjectStudyTime, getAggregationSubjectStudyTimeQuery)
	if err != nil {
		log.Printf("GetAggregationSubjectStudyTime:%s\n", err)
	}
}
