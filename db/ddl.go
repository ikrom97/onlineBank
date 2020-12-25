package db

const (
	CreateUserTable = `create table if not exists user (
	id integer primary key autoincrement unique,
	name text not null,
	surname text not null,
	age integer not null,
	gender text not null,
	admin boolean not null,
	login text not null unique,
	password text not null,
	remove boolean default false)`

	CreateAccountTable = `create table if not exists account (
	id integer primary key autoincrement unique,
	userId integer references user(id),
	name text not null,
	number integer not null unique,
	amount integer not null,
	currency text not null)`

	CreateATMTable = `create table if not exists atm (
	id integer primary key autoincrement,
	address text not null,
	status boolean not null default true)`

	CreateTransactionTable = `create table if not exists transactionHistory (
	id integer primary key autoincrement,
	date text not null,
	time text not null,
	operationAmount integer not null,
	accountNumber integer not null,
	receiverAccountNumber integer not null,
	availableLimit integer not null)`
)
