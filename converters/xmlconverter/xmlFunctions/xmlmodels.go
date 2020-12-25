package xmlFunctions

import "encoding/xml"

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}
type User struct {
	XMLName  xml.Name `xml:"user"`
	ID       int64    `xml:"id"`
	Name     string   `xml:"name"`
	Surname  string   `xml:"surname"`
	Age      int64    `xml:"age"`
	Gender   string   `xml:"gender"`
	Admin    bool     `xml:"admin"`
	Login    string   `xml:"login"`
	Password string   `xml:"password"`
	Remove   bool     `xml:"remove"`
}
type Accounts struct {
	XMLName  xml.Name  `xml:"accounts"`
	Accounts []Account `xml:"accounts"`
}
type Account struct {
	XMLName  xml.Name `xml:"account"`
	ID       int64    `xml:"id"`
	UserID   int64    `xml:"user_id"`
	Name     string   `xml:"name"`
	Number   int64    `xml:"number"`
	Amount   int64    `xml:"amount"`
	Currency string   `xml:"currency"`
}
type ATMs struct {
	XMLName xml.Name `xml:"atms"`
	ATMs    []ATM    `xml:"atm"`
}
type ATM struct {
	XMLName xml.Name `xml:"atm"`
	ID      int64    `xml:"id"`
	Address string   `xml:"address"`
	Status  bool     `xml:"status"`
}
