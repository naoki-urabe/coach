package models

import (
	"crypto/sha256"
	"fmt"
	"log"
)

type User struct {
	Id         string `db:"id" json:"id"`
	Pw         string `db:"pw" json:"pw"`
	PrivateKey string `db:"private_key" json:"private_key"`
	PublicKey  string `db:"public_key" json:"public_key"`
}

var insertUserQuery = `
INSERT INTO users VALUES(?,?,?,?);
`
var findUser = `
SELECT * FROM users WHERE id = ? AND pw = ? LIMIT 1;`

var checkDuplicateUserQuery = `
SELECT * FROM users WHERE id = ?`

func InsertUser(user *User) bool {
	_, err := Db.Queryx(insertUserQuery, user.Id, user.Pw, user.PrivateKey, user.PublicKey)
	//errがnilだったら登録失敗
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func FindUser(user *User) bool {
	p := []byte(user.Pw)
	sha256 := sha256.Sum256(p)
	user.Pw = fmt.Sprintf("%x", sha256)
	err := Db.Get(user, findUser, user.Id, user.Pw)
	//errがnilじゃなかったらログインできない
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
