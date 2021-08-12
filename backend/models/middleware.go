package models

import (
	"coach/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Subject struct {
	SubjectCode string `json:"subject_code"`
	Subject     string `json:"subject"`
}

type User struct {
	Id         string `json:"id"`
	Password   string `json:"password"`
	PrivateKey string `gorm:"size:2048" json:"private_key"`
	PublicKey  string `gorm:"size:2048" json:"public_key"`
}

var Db *gorm.DB

func GetAllSubject(subject *[]Subject) {
	Db.Find(&subject)
}

func InsertSubject(subject *Subject) {
	Db.NewRecord(subject)
	Db.Create(&subject)
}

func InsertUser(user *User) {
	Db.NewRecord(user)
	Db.Create(&user)
}

func FindUser(user *User) {
	Db.Where("Id = ?", user.Id).Where("PASSWORD = ?", user.Password).Find(&user)
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
	Db, err = gorm.Open(config.Config.DbDriverName, dbConnectInfo)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connect database..")
	}
	Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Subject{})
	Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created table...")
	}
}
