package services

import (
	"database/sql"
	"fmt"
	"log"
	"onlineBank/models"
)

const AdminOperationWindow = `1.Показать баланс
2.Переводы
3.Адреса банкоматов
4.Добавить банкомат
5.История транзакций
0.Назад`

func AdminOperation(Db *sql.DB, user models.User) {
	fmt.Printf("Привет %s, выбери команду...\n", user.Name)
	fmt.Println(AdminOperationWindow)
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Println("Неверный ввод, введите число")
		return
	}
	switch input {
	case 1:
		ShowAccountAmount(Db, user)
	case 2:
		Translations(Db, user)
	case 3:
		ShowATMsAddresses(Db, user)
	case 4:
		AddNewATM(Db, user)
	case 5:
		ShowTransactions(Db, user)
	case 0:
		AuthorizeAndStart(Db)
	default:
		fmt.Println("Выберите от 1 до 6")
	}
}
