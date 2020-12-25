package models

import (
	"database/sql"
	"onlineBank/db"
	"time"
)

type Transaction struct {
	ID                    int64
	Date                  string
	Time                  string
	OperationAmount       int64
	AccountNumber         int64
	ReceiverAccountNumber int64
	AvailableLimit        int64
}

func AddTransaction(Db *sql.DB, myAccount, receiverAccount Account, operationAmount int64) (err error) {
	var check Transaction
	data := time.Now()
	check.Date = data.Format("02-Jan-2006")
	check.Time = data.Format("15:40")
	check.OperationAmount = operationAmount
	check.AccountNumber = myAccount.Number
	check.AvailableLimit = myAccount.Amount - operationAmount
	check.ReceiverAccountNumber = receiverAccount.Number
	_, err = Db.Exec(db.AddTransaction, check.Date, check.Time, check.OperationAmount, check.AccountNumber, check.ReceiverAccountNumber, check.AvailableLimit)
	if err != nil {
		return err
	}
	return
}
