package models

import (
	"crypto/sha256"
	"fmt"
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

func InsertUser(user *User) {
	Db.MustExec(insertUserQuery, user.Id, user.Pw, user.PrivateKey, user.PublicKey)
}

func FindUser(user *User) {
	p := []byte(user.Pw)
	sha256 := sha256.Sum256(p)
	user.Pw = fmt.Sprintf("%x", sha256)
	Db.Get(user, findUser, user.Id, user.Pw)
}
