package services

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"onlineBank/db"
	"onlineBank/models"
	"os"
)

func ShowAccountAmount(Db *sql.DB, user models.User) {
	rows, err := Db.Query(db.GetAccountByUserId, user.ID)
	if err != nil {
		log.Println("Cannot get account by id")
		return
	}
	var accounts []models.Account
	for rows.Next() {
		p := models.Account{}
		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Name,
			&p.Number,
			&p.Amount,
			&p.Currency)
		if err != nil {
			fmt.Println(err)
			continue
		}
		accounts = append(accounts, p)
	}
	if len(accounts) > 1 {
		fmt.Println("Ваш баланс на карте...")
		var total int64
		for _, p := range accounts {
			fmt.Printf("%s составляет %d %s \n", p.Name, p.Amount, p.Currency)
			total += p.Amount
		}
		fmt.Printf("В сумме у Вас %d %s", total, accounts[0].Currency)
	} else if len(accounts) == 1 {
		fmt.Printf("Ваш баланс на карте %s составляет %d %s \n", accounts[0].Name, accounts[0].Amount, accounts[0].Currency)
	} else {
		fmt.Println("Пока что у Вас нет аккаунта")
	}
	UserAdminExit(Db, user)
}
func Translations(Db *sql.DB, user models.User) {
	var myAccountNumber, receiverAccountNumber, translationAmount int64
	fmt.Println("Введите номер вашей карты:")
	_, err := fmt.Scan(&myAccountNumber)
	if err != nil {
		log.Print("Неверный ввод")
		return
	}
	fmt.Println("Введите номер карты получателя:")
	_, err = fmt.Scan(&receiverAccountNumber)
	if err != nil {
		log.Print("Неверный ввод")
		return
	}
	fmt.Println("Сумма:")
	_, err = fmt.Scan(&translationAmount)
	if err != nil {
		log.Print("Неверный ввод")
		return
	}
	var myAccount, receiverAccount models.Account
	row := Db.QueryRow(db.GetAccountAmount, myAccountNumber)
	err = row.Scan(&myAccount.Amount)
	if err != nil {
		log.Println(err)
		return
	}
	row = Db.QueryRow(db.GetAccountNumber, myAccountNumber)
	err = row.Scan(&myAccount.Number)
	if err != nil {
		log.Println(err)
		return
	}
	row = Db.QueryRow(db.GetAccountAmount, receiverAccountNumber)
	err = row.Scan(&receiverAccount.Amount)
	if err != nil {
		log.Println(err)
		return
	}
	row = Db.QueryRow(db.GetAccountNumber, receiverAccountNumber)
	err = row.Scan(&receiverAccount.Number)
	if err != nil {
		log.Println(err)
		return
	}
	if myAccount.Amount < translationAmount {
		fmt.Println("У Вас недостаточно средств!")
		UserAdminExit(Db, user)
		return
	}
	err = TransferToAccount(Db, myAccountNumber, receiverAccountNumber, translationAmount)
	if err != nil {
		log.Println(err)
		return
	}
	err = models.AddTransaction(Db, myAccount, receiverAccount, translationAmount)
	if err != nil {
		log.Println(err)
		return
	}
	UserAdminExit(Db, user)
}
func TransferToAccount(Db *sql.DB, giverAccountNumber, gainerAccountNumber, amount int64) (err error) {
	tx, err := Db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	_, err = Db.Exec(db.UpdateAccountAmountOfGiver, amount, giverAccountNumber)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = Db.Exec(db.UpdateAccountAmountOfGainer, amount, gainerAccountNumber)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Ваша сумма успешно переведено!")
	return
}
func ShowATMsAddresses(Db *sql.DB, user models.User) {
	rows, err := Db.Query(db.GetATM)
	if err != nil {
		log.Println(err)
		return
	}
	atms := []models.ATM{}
	for rows.Next() {
		p := models.ATM{}
		err := rows.Scan(
			&p.ID,
			&p.Address,
			&p.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		atms = append(atms, p)
	}
	for _, p := range atms {
		fmt.Println(p)
	}
	UserAdminExit(Db, user)
}
func ShowTransactions(Db *sql.DB, user models.User) {
	rows, err := Db.Query(db.GetAccountByUserId, user.ID)
	if err != nil {
		log.Println(err)
		return
	}
	acc := []models.Account{}
	for rows.Next() {
		p := models.Account{}
		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Name,
			&p.Number,
			&p.Amount,
			&p.Currency)
		if err != nil {
			fmt.Println(err)
			continue
		}
		acc = append(acc, p)
	}
	var t int
	for i := 0; i < len(acc); i++ {
		t = i
	}
	if t < 1 {
		rows, err := Db.Query(db.GetTransaction, acc[0].Number)
		if err != nil {
			log.Println(err)
			return
		}
		arr := []models.Transaction{}
		for rows.Next() {
			p := models.Transaction{}
			err := rows.Scan(
				&p.ID,
				&p.Date,
				&p.Time,
				&p.OperationAmount,
				&p.AccountNumber,
				&p.ReceiverAccountNumber,
				&p.AvailableLimit)
			if err != nil {
				fmt.Println(err)
				continue
			}
			arr = append(arr, p)
		}
		for i := 0; i < len(arr); i++ {
			fmt.Println("    ОАО ХУМО БАНК   ")
			fmt.Println("Номер банкомата:", arr[i].ID)
			fmt.Println("Дата:", arr[i].Date)
			fmt.Println("Время:", arr[i].Time)
			fmt.Println("Сумма операции:", arr[i].OperationAmount)
			fmt.Println("Номер вашей карты:", arr[i].AccountNumber)
			fmt.Println("Номер карты получателя:", arr[i].ReceiverAccountNumber)
			fmt.Println("Доступный лимит:", arr[i].AvailableLimit, "\n")
		}
	} else {
		fmt.Println("Введите номер вашей карты:")
		var accountNumber int64
		_, err := fmt.Scan(&accountNumber)
		if err != nil {
			log.Println(err)
			return
		}
		rows, err := Db.Query(db.GetTransaction, accountNumber)
		if err != nil {
			log.Println(err)
			return
		}
		arr := []models.Transaction{}
		for rows.Next() {
			p := models.Transaction{}
			err := rows.Scan(
				&p.ID,
				&p.Date,
				&p.Time,
				&p.OperationAmount,
				&p.AccountNumber,
				&p.ReceiverAccountNumber,
				&p.AvailableLimit)
			if err != nil {
				fmt.Println(err)
				continue
			}
			arr = append(arr, p)
		}
		for i := 0; i < len(arr); i++ {
			fmt.Println("    ОАО ХУМО БАНК   ")
			fmt.Println("Номер банкомата:", arr[i].ID)
			fmt.Println("Дата:", arr[i].Date)
			fmt.Println("Время:", arr[i].Time)
			fmt.Println("Сумма операции:", arr[i].OperationAmount)
			fmt.Println("Номер вашей карты:", arr[i].AccountNumber)
			fmt.Println("Номер карты получателя:", arr[i].ReceiverAccountNumber)
			fmt.Println("Доступный лимит:", arr[i].AvailableLimit, "\n")
		}
	}
	UserAdminExit(Db, user)
}
func UserAdminExit(db *sql.DB, user models.User) {
	fmt.Println("\n 0.Назад")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Println("Неверный ввод", err)
	}
	switch input {
	case 0:
		if user.Admin == true {
			AdminOperation(db, user)
		} else {
			UserOperation(db, user)
		}
	default:
		log.Println("Неверный ввод", err)
	}
}
func AddNewATM(Db *sql.DB, user models.User) {
	fmt.Println("Enter your address")
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Println(err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	address, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(s)
	sprint := fmt.Sprintf("%s %s", s, address)
	fmt.Println(sprint)
	err = models.AddNewATM(Db, sprint)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Ваш банкомат успешно добавлен")
	UserAdminExit(Db, user)
}
