package services

import (
	"database/sql"
	"fmt"
	"log"
	"onlineBank/models"
	"os"
)

const UserOperationWindow = `1.Показать баланс
2.Переводы
3.Адреса банкоматов
4.История транзакций
0.Выйти`

func UserOperation(Db *sql.DB, user models.User) {
	fmt.Printf("Привет %s, выбери команду...\n", user.Name)
	fmt.Println(UserOperationWindow)
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Println("Неверный ввод, введите число", err)
	}
	switch input {
	case 1:
		ShowAccountAmount(Db, user)
	case 2:
		Translations(Db, user)
	case 3:
		ShowATMsAddresses(Db, user)
	case 4:
		ShowTransactions(Db, user)
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Неверный ввод")
	}
}
