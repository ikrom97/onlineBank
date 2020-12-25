package models

import (
	"database/sql"
	"onlineBank/db"
)

type ATM struct {
	ID      int64  `json:"id" xml:"id"`
	Address string `json:"address" xml:"address"`
	Status  bool   `json:"status" xml:"status"`
}

func (receiver ATM) AddNewATM(Db *sql.DB, atm ATM) (err error) {
	_, err = Db.Exec(db.AddATM, atm.Address)
	if err != nil {
		return err
	}
	return
}
func AddNewATM(Db *sql.DB, address string) (err error) {
	_, err = Db.Exec(db.AddNewATM, address)
	if err != nil {
		return err
	}
	return
}
func AddATM(Db *sql.DB, atm ATM) (err error) {
	_, err = Db.Exec(db.AddATM, atm.ID, atm.Address, atm.Status)
	if err != nil {
		return err
	}
	return
}
