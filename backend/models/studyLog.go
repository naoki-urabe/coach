package models

import (
	_ "fmt"
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

func AddStudyStartLog(startLog *StartLog) int {
	registerLog := Db.MustExec(addStudyStartLogQuery, startLog.SubjectCode, startLog.StudyStartTime, startLog.User)
	id, _ := registerLog.LastInsertId()
	return int(id)
}

func AddStudyFinishLog(finishLog *FinishLog) {
	Db.MustExec(addStudyFinishLogQuery, finishLog.Content, finishLog.Comment, finishLog.StudyFinishTime, finishLog.Id)
}

func GetAllStudyLog(user string, studyLog *[]StudyLog) {
	Db.Select(studyLog, getAllStudyLogQuery, user)
}

func GetSubjectStudyLog(user string, subjectCode string, studyLog *[]StudyLog) {
	Db.Select(studyLog, getSubjectStudyLogQuery, user, subjectCode)
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
