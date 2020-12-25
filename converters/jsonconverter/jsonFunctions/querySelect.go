package jsonFunctions

import (
	"database/sql"
	"log"
	"onlineBank/db"
	"onlineBank/models"
)

func GetAccountsFromDB(Db *sql.DB) (accounts []models.Account, err error) {
	row, err := Db.Query(db.GetAccount)
	if err != nil {
		log.Println("Can't get account", err)
		return
	}
	for row.Next() {
		p := models.Account{}
		err := row.Scan(&p.ID, &p.UserID, &p.Name, &p.Number, &p.Amount, &p.Currency)
		if err != nil {
			log.Println(err)
			continue
		}
		accounts = append(accounts, p)
	}
	return
}
func GetATMsFromDB(Db *sql.DB) (atms []models.ATM, err error) {
	row, err := Db.Query(db.GetATM)
	if err != nil {
		log.Println(err)
		return
	}
	for row.Next() {
		p := models.ATM{}
		err := row.Scan(&p.ID, &p.Address, &p.Status)
		if err != nil {
			log.Println(err)
			continue
		}
		atms = append(atms, p)
	}
	return
}
func GetUsersFromDB(Db *sql.DB) (users []models.User, err error) {
	row, err := Db.Query(db.GetUser)
	if err != nil {
		log.Println(err)
		return
	}
	for row.Next() {
		p := models.User{}
		err := row.Scan(&p.ID, &p.Name, &p.Surname, &p.Age, &p.Gender, &p.Admin, &p.Login, &p.Password, &p.Remove)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, p)
	}
	return
}
