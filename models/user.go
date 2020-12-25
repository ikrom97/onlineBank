package models

import (
	"database/sql"
	"onlineBank/db"
)

type User struct {
	ID       int64  `json:"id" xml:"id"`
	Name     string `json:"name" xml:"name"`
	Surname  string `json:"surname" xml:"surname"`
	Age      int64  `json:"age" xml:"age"`
	Gender   string `json:"gender" xml:"gender"`
	Admin    bool   `json:"admin" xml:"admin"`
	Login    string `json:"login" xml:"login"`
	Password string `json:"password" xml:"password"`
	Remove   bool   `json:"remove" xml:"remove"`
}

func (receiver User) AddNewUser(Db *sql.DB, user User) (err error) {
	_, err = Db.Exec(db.AddUser, user.Name, user.Surname, user.Age, user.Gender, user.Admin, user.Login, user.Password)
	if err != nil {
		return err
	}
	return
}
func AddNewUser(Db *sql.DB, user User) (err error) {
	_, err = Db.Exec(db.AddUser, user.Name, user.Surname, user.Age, user.Gender, user.Admin, user.Login, user.Password)
	if err != nil {
		return err
	}
	return
}
