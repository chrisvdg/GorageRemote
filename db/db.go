package db

import (
	"database/sql"
	"errors"
	"fmt"

	e "github.com/chrisvdg/GorageRemote/entities"
	"golang.org/x/crypto/bcrypt"
	// import sqlite3 capabilities
	_ "github.com/mattn/go-sqlite3"
)

var (
	// ErrDBConn is an error when something went wrong with the database connection
	ErrDBConn = errors.New("something went wrong with the database connection")
	// ErrUserNotFound is an error where the requested used was not found in the database
	ErrUserNotFound = errors.New("user not found in database")
	// ErrFailedAuth is an error where the user could not be authenticated
	ErrFailedAuth = errors.New("failed to authenticate user")
)

// NewDB returns db connections from provided sqlite source file path
// also initialises db with neccesary tables (and default admin)
func NewDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil || db == nil {
		return nil, ErrDBConn
	}

	err = initDB(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// InitDB checks if database exists
// if not it will create it with needed tables and fields
func initDB(db *sql.DB) error {
	// check if tables are present
	//users
	checkTableUsers := `
SELECT name FROM sqlite_master WHERE type='table' AND name=?;
	`
	var exists string
	err := db.QueryRow(checkTableUsers, "users").Scan(&exists)
	if err != sql.ErrNoRows {
		return err
	}
	// create table
	tableUsers := `
CREATE TABLE IF NOT EXISTS users(
	username TEXT NOT NULL PRIMARY KEY,
	password TEXT NOT NULL,
	admin BOOL NOT NULL
);
`
	_, err = db.Exec(tableUsers)
	if err != nil {
		return err
	}

	u := e.User{
		Name:     "admin",
		Password: "Gorage123",
		Admin:    true,
	}

	return AddUser(db, u)
}

// CheckPassword checks password of provided user
func CheckPassword(db *sql.DB, u *e.User) error {
	var passBuf string
	q := "SELECT password FROM users WHERE username is ?;"
	err := db.QueryRow(q, u.Name).Scan(&passBuf)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passBuf), []byte(u.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return ErrFailedAuth
		}
		return err
	}

	return nil
}

// UpdateAdmin updates a user's admin rights
func UpdateAdmin(db *sql.DB, u *e.User) error {
	if u.Name == "" {
		return fmt.Errorf("no username was provided")
	}

	q := "UPDATE users SET admin=? WHERE username=?;"
	_, err := db.Exec(q, u.Admin, u.Name)

	return err
}

// UpdatePassword updates a user's password
func UpdatePassword(db *sql.DB, u *e.User) error {
	if u.Name == "" {
		return fmt.Errorf("no username was provided")
	}
	if u.Password == "" {
		return fmt.Errorf("no password was provided")
	}

	q := "UPDATE users SET password=? WHERE username=?;"
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	_, err := db.Exec(q, hash, u.Name)

	return err
}

// GetUser fetches a used from the database
func GetUser(db *sql.DB, uname string) (*e.User, error) {
	if uname == "" {
		return nil, fmt.Errorf("no username was provided")
	}
	u := new(e.User)
	var adminBuf bool

	q := "SELECT admin FROM users WHERE username is ?;"
	err := db.QueryRow(q, uname).Scan(&adminBuf)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	u.Name = uname
	u.Admin = adminBuf

	return u, nil
}

// AddUser adds user to the database
// encrypts password before adding
func AddUser(db *sql.DB, user e.User) error {
	if user.Name == "" {
		return fmt.Errorf("no username was provided")
	}
	if user.Password == "" {
		return fmt.Errorf("no password was provided")
	}

	addAdmin := `
	INSERT INTO users(
		username,
		password,
		admin
	)VALUES(?,?,?);
			`
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	_, err := db.Exec(addAdmin, "admin", string(hash), true)
	if err != nil {
		return err
	}
	return nil
}
