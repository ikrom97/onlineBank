package xmlFunctions

import (
	"database/sql"
	"encoding/xml"
	"io/ioutil"
	"log"
	"onlineBank/converters/jsonconverter/jsonFunctions"
	"onlineBank/models"
	"os"
)

func AddDBToXML(Db *sql.DB) {
	err := AddUsersToXML(Db)
	if err != nil {
		log.Println(err)
		return
	}
	err = AddAccountsToXML(Db)
	if err != nil {
		log.Println(err)
		return
	}
	err = AddATMsToXML(Db)
	if err != nil {
		log.Println(err)
		return
	}
}
func AddUsersToXML(Db *sql.DB) (err error) {
	var users Users
	user, _ := jsonFunctions.GetUsersFromDB(Db)
	for i := 0; i < len(user); i++ {
		usr := User{
			XMLName:  xml.Name{},
			ID:       user[i].ID,
			Name:     user[i].Name,
			Surname:  user[i].Surname,
			Age:      user[i].Age,
			Gender:   user[i].Gender,
			Admin:    user[i].Admin,
			Login:    user[i].Login,
			Password: user[i].Password,
			Remove:   user[i].Remove,
		}
		users.Users = append(users.Users, usr)
	}
	xmlUsers, _ := xml.Marshal(&users)
	file, err := os.Create("converters/xmlconverter/users.xml")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = file.Write(xmlUsers)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
func AddAccountsToXML(Db *sql.DB) (err error) {
	var accounts Accounts
	account, _ := jsonFunctions.GetAccountsFromDB(Db)
	for i := 0; i < len(account); i++ {
		acc := Account{
			XMLName:  xml.Name{},
			ID:       account[i].ID,
			UserID:   account[i].UserID,
			Name:     account[i].Name,
			Number:   account[i].Number,
			Amount:   account[i].Amount,
			Currency: account[i].Currency,
		}
		accounts.Accounts = append(accounts.Accounts, acc)
	}
	xmlAccounts, _ := xml.Marshal(&accounts)
	file, err := os.Create("converters/xmlconverter/accounts.xml")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = file.Write(xmlAccounts)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
func AddATMsToXML(Db *sql.DB) (err error) {
	var atms ATMs
	atm, _ := jsonFunctions.GetATMsFromDB(Db)
	for i := 0; i < len(atm); i++ {
		at := ATM{
			XMLName: xml.Name{},
			ID:      atm[i].ID,
			Address: atm[i].Address,
			Status:  atm[i].Status,
		}
		atms.ATMs = append(atms.ATMs, at)
	}
	xmlATMs, _ := xml.Marshal(&atms)
	file, err := os.Create("converters/xmlconverter/atms.xml")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = file.Write(xmlATMs)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
func AddXMLToDB(Db *sql.DB) {
	err := AddXMLUsersToDB(Db)
	if err != nil {
		log.Println(err)
		return
	}
	err = AddXMLAccountsToDB(Db)
	if err != nil {
		log.Println(err)
		return
	}
	err = AddXMLATMsToDB(Db)
	if err != nil {
		log.Println(err)
		return
	}
}
func AddXMLUsersToDB(Db *sql.DB) (err error) {
	xmlUsers, err := ioutil.ReadFile("converters/xmlconverter/users.xml")
	if err != nil {
		log.Println(err)
		return
	}
	var users Users
	err = xml.Unmarshal(xmlUsers, &users)
	if err != nil {
		log.Println(err)
		return
	}
	var usr []models.User
	for i := 0; i < len(users.Users); i++ {
		p := models.User{
			ID:       users.Users[i].ID,
			Name:     users.Users[i].Name,
			Surname:  users.Users[i].Surname,
			Age:      users.Users[i].Age,
			Gender:   users.Users[i].Gender,
			Admin:    users.Users[i].Admin,
			Login:    users.Users[i].Login,
			Password: users.Users[i].Password,
			Remove:   users.Users[i].Remove,
		}
		usr = append(usr, p)
	}
	for _, value := range usr {
		_ = models.AddNewUser(Db, value)
	}
	return
}
func AddXMLAccountsToDB(Db *sql.DB) (err error) {
	xmlAccounts, err := ioutil.ReadFile("converters/xmlconverter/accounts.xml")
	if err != nil {
		log.Println(err)
		return
	}
	var accounts Accounts
	err = xml.Unmarshal(xmlAccounts, &accounts)
	if err != nil {
		log.Println(err)
		return
	}
	var acc []models.Account
	for i := 0; i < len(accounts.Accounts); i++ {
		p := models.Account{
			ID:       accounts.Accounts[i].ID,
			UserID:   accounts.Accounts[i].UserID,
			Name:     accounts.Accounts[i].Name,
			Number:   accounts.Accounts[i].Number,
			Amount:   accounts.Accounts[i].Amount,
			Currency: accounts.Accounts[i].Currency,
		}
		acc = append(acc, p)
	}
	for _, value := range acc {
		_ = models.AddNewAccount(Db, value)
	}
	return
}
func AddXMLATMsToDB(Db *sql.DB) (err error) {
	xmlATMs, err := ioutil.ReadFile("converters/xmlconverter/atms.xml")
	if err != nil {
		log.Println(err)
		return
	}
	var atms ATMs
	err = xml.Unmarshal(xmlATMs, &atms)
	if err != nil {
		log.Println(err)
		return
	}
	var at []models.ATM
	for i := 0; i < len(atms.ATMs); i++ {
		p := models.ATM{
			ID:      atms.ATMs[i].ID,
			Address: atms.ATMs[i].Address,
			Status:  atms.ATMs[i].Status,
		}
		at = append(at, p)
	}
	for _, value := range at {
		_ = models.AddATM(Db, value)
	}
	return
}
