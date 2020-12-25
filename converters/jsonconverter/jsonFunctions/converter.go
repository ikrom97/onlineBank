package jsonFunctions

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"onlineBank/models"
	"os"
)

func AddDBToJson(Db *sql.DB) {
	accounts, _ := GetAccountsFromDB(Db)
	acc, _ := json.Marshal(accounts)
	file, err := os.Create("converters/jsonconverter/accounts.json")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = file.Write(acc)
	if err != nil {
		log.Println(err)
		return
	}
	atms, _ := GetATMsFromDB(Db)
	atm, _ := json.Marshal(atms)
	file1, err := os.Create("converters/jsonconverter/atms.json")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = file1.Write(atm)
	if err != nil {
		log.Println(err)
		return
	}

	users, _ := GetUsersFromDB(Db)
	user, _ := json.Marshal(users)
	file2, err := os.Create("converters/jsonconverter/users.json")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = file2.Write(user)
	if err != nil {
		log.Println(err)
		return
	}
}
func AddJsonToDB(Db *sql.DB) {
	var accounts []models.Account
	file, err := ioutil.ReadFile("converters/jsonconverter/accounts.json")
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(file, &accounts)
	if err != nil {
		log.Println(err)
		return
	}
	for _, value := range accounts {
		_ = models.AddNewAccount(Db, value)
	}
	var atms []models.ATM
	file1, err := ioutil.ReadFile("converters/jsonconverter/atms.json")
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(file1, &atms)
	if err != nil {
		log.Println(err)
		return
	}
	for _, value := range atms {
		_ = models.AddATM(Db, value)
	}
	var users []models.User
	file2, err := ioutil.ReadFile("converters/jsonconverter/users.json")
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(file2, &users)
	if err != nil {
		log.Println(err)
		return
	}
	for _, value := range users {
		_ = models.AddNewUser(Db, value)
	}

}
