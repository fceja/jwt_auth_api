package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

type Storage interface {
	createAccount(*Account) error
	deleteAccount(int) error
	getAccounts() ([]*Account, error)
	getAccountByID(int) (*Account, error)
	getAccountByNumber(int) (*Account, error)
	updateAccount(*Account) error
}

func newPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=letMeIn2020 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	acnt := new(Account)
	err := rows.Scan(
		&acnt.ID,
		&acnt.FirstName,
		&acnt.LastName,
		&acnt.Number,
		&acnt.CreatedAt)

	return acnt, err
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query :=
		`create table if not exists account(
			id serial primary key,
			first_name varchar(50),
			last_name varchar(50),
			number serial,
			encrypted_password varchar(50),
			balance serial,
			created_at timestamp
		)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) createAccount(acnt *Account) error {
	query :=
		`insert into account(
			first_name,
			last_name,
			number,
			encrypted_password,
			balance,
			created_at)
		values (
			$1, $2, $3, $4, $5, $6
		)`

	_, err := s.db.Query(
		query,
		acnt.FirstName,
		acnt.LastName,
		acnt.Number,
		acnt.EncryptedPassword,
		acnt.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) updateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) deleteAccount(id int) error {
	_, err := s.db.Query("delete from account where id = $1", id)
	return err
}

func (s *PostgresStore) getAccountByNumber(num int) (*Account, error) {
	rows, err := s.db.Query("select * from account where number = $1", num)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account with number [%d] not found", num)
}

func (s *PostgresStore) getAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query("select * from account where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) getAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		acnt := new(Account)
		err := rows.Scan(
			&acnt.ID,
			&acnt.FirstName,
			&acnt.LastName,
			&acnt.Number,
			&acnt.CreatedAt)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, acnt)
	}

	return accounts, nil
}
