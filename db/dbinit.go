package db

import "database/sql"

func DatabaseInit(db *sql.DB) (err error) {
	DDLs := []string{CreateUserTable, CreateAccountTable, CreateATMTable, CreateTransactionTable}
	for _, ddl := range DDLs {
		_, err = db.Exec(ddl)
		if err != nil {
			return err
		}
	}
	return
}
