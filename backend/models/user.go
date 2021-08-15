package models

type User struct {
	Id         string `db:"id" json:"id"`
	Password   string `db:"password" json:"password"`
	PrivateKey string `db:"private_key" json:"private_key"`
	PublicKey  string `db:"public_key" json:"public_key"`
}

var insertUserQuery = `
INSERT INTO users VALUES(?,?,?,?);
`
var findUser = `
SELECT * FROM users WHERE id = ? AND password = ? LIMIT 1;`

func InsertUser(user *User) {
	Db.MustExec(insertUserQuery, user.Id, user.Password, user.PrivateKey, user.PublicKey)
}

func FindUser(user *User) {
	Db.Get(user, findUser, user.Id, user.Password)
}