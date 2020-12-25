package models

import (
	"database/sql"
	"onlineBank/db"
)

type Account struct {
	ID       int64  `json:"id" xml:"id"`
	UserID   int64  `json:"user_id" xml:"user_id"`
	Name     string `json:"name" xml:"name"`
	Number   int64  `json:"number" xml:"number"`
	Amount   int64  `json:"amount" xml:"amount"`
	Currency string `json:"currency" xml:"currency"`
}

func (receiver Account) AddNewAccount(Db *sql.DB, account Account) (err error) {
	_, err = Db.Exec(db.AddNewAccount, account.UserID, account.Name, account.Number, account.Amount, account.Currency)
	if err != nil {
		return err
	}
	return nil
}
func AddNewAccount(Db *sql.DB, account Account) (err error) {
	_, err = Db.Exec(db.AddNewAccount, account.UserID, account.Name, account.Number, account.Amount, account.Currency)
	if err != nil {
		return err
	}
	return nil
}
