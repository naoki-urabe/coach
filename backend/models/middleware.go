package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"coach/config"
)

type Subject struct {
	SubjectCode string `json:"subject_code"`
	Subject     string `json:"subject"`
}

var Db *gorm.DB

func GetAllSubject(subject *[]Subject) {
	Db.Find(&subject)
}

func InsertSubject(subject *Subject) {
	Db.NewRecord(subject)
	Db.Create(&subject)
}

func init() {
	var err error
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		config.Config.DbUserName,
		config.Config.DbUserPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbName,
	)
	fmt.Println(dbConnectInfo)
	Db, err = gorm.Open(config.Config.DbDriverName, dbConnectInfo)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connect database..")
	}
	Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Subject{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created table...")
	}
}
