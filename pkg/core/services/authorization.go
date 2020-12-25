package services

import (
	"database/sql"
	"fmt"
	"log"
	"onlineBank/db"
	"onlineBank/models"
	"os"
)

const AuthorizationWindow = `1.Авторизация
2.Выйти`

func AuthorizeAndStart(Db *sql.DB) {
	fmt.Println("Выберите команду...")
	fmt.Println(AuthorizationWindow)
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Println("Неверный ввод, введите число")
		return
	}
	switch input {
	case 1:
		login, password := AuthorizationOperation(Db)
		user, err := GetUserByLoginAndPassword(Db, login, password)
		if err != nil {
			log.Println(err)
			return
		}
		if user.Admin {
			AdminOperation(Db, user)
		} else {
			UserOperation(Db, user)
		}
	case 2:
		os.Exit(0)
	default:
		log.Println("Выберите 1 или 2")
		return
	}
}
func AuthorizationOperation(Db *sql.DB) (login, password string) {
	fmt.Println("Введите Ваш логин и пароль...\nЛогин:")
	_, err := fmt.Scan(&login)
	if err != nil {
		log.Println("Неверный логин")
		return
	}
	fmt.Println("Пароль:")
	_, err = fmt.Scan(&password)
	if err != nil {
		log.Println("Неверный пароль", err)
		return
	}
	return
}
func GetUserByLoginAndPassword(Db *sql.DB, login, password string) (user models.User, err error) {
	row := Db.QueryRow(db.GetUserByLoginPassword, login, password)
	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Age,
		&user.Gender,
		&user.Admin,
		&user.Login,
		&user.Password,
		&user.Remove)
	if err != nil {
		log.Println("Извините но вы не зарегистрированы")
		return
	}
	return
}
