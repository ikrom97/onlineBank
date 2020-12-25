package db

const (
	AddNewAccount = `insert into account(userId, name, number, amount, currency)
values(($1),($2),($3),($4),($6))`
	AddNewATM      = `insert into atm(address) values(($1))`
	AddATM         = `insert into atm(id, address, status) values(($1),($2),($3))`
	AddTransaction = `insert into transaction(date, time, operationAmount, accountNumber, receiverAccountNumber, availableLimit) 
values(($1),($2),($3),($4),($5),($6))`
	AddUser = `insert into user(name, surname, age, gender, admin, login, password) 
values(($1),($2),($3),($4),($5),($6),($7))`
	GetUserByLoginPassword      = `select *from user where login = ($1) and password = ($2)`
	GetAccountByUserId          = `select * from account where account.userId = ($1)`
	GetAccountAmount            = `select amount from account where number = ($1)`
	GetAccountNumber            = `select number from account where number = ($1)`
	UpdateAccountAmountOfGiver  = `update account set amount = amount - ($1) where number = ($2)`
	UpdateAccountAmountOfGainer = `update account set amount = amount + ($1) where number = ($2)`
	GetATM                      = `select * from atm`
	GetAccount                  = `select * from account`
	GetUser                     = `select * from user`
	GetTransaction              = `select * from transaction where accountNumber = ($1)`
)
