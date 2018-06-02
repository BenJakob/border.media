package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error
var dbDateLayout = "20060102"
var localDateLayout = "02.01.2006"
var localDateLayoutInputField = "2006-01-02"

func init() {
	db, err = sql.Open("sqlite3", "app/model/Web.db")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
