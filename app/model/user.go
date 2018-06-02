package model

import (
	"net/http"
	"time"
)

type User struct {
	ID       uint32
	Name     string
	LastName string
	UserName string
	Email    string
	Created  string
	Password string
	Image    string
	Status   string
	Orders   []LedgerEntry
}

func GetUser(id uint32) (User, error) {
	user := User{}
	err = db.QueryRow(`
		SELECT User.ID, User.Name, [Last Name], [User Name], Created, Email, Password, Image, [User Status].Name
		FROM User
		INNER JOIN [User Status] on User.Status = [User Status].ID
		WHERE User.ID = $1`, id).Scan(&user.ID, &user.Name, &user.LastName, &user.UserName, &user.Created, &user.Email, &user.Password, &user.Image, &user.Status)
	parsedDate, _ := time.Parse(dbDateLayout, user.Created)
	user.Created = parsedDate.Format(localDateLayout)
	return user, err
}

func GetUserByName(userName string) (User, error) {
	user := User{}
	err = db.QueryRow(`
		SELECT User.ID, User.Name, [Last Name], [User Name], Created, Email, Password, Image, [User Status].Name
		FROM User
		INNER JOIN [User Status] on User.Status = [User Status].ID
		WHERE [User Name] = $1`, userName).Scan(&user.ID, &user.Name, &user.LastName, &user.UserName, &user.Created, &user.Email, &user.Password, &user.Image, &user.Status)
	parsedDate, _ := time.Parse(dbDateLayout, user.Created)
	user.Created = parsedDate.Format(localDateLayout)
	return user, err
}

func GetLoggedInUser(w http.ResponseWriter, r *http.Request) (User, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return User{}, err
	}

	session, err := GetSession(cookie.Value)
	if err != nil {
		return User{}, err
	}

	return GetUser(session.UserID)
}

func GetUsers() []User {
	rows, err := db.Query(`
		SELECT User.ID, User.Name, [Last Name], [User Name], Email, Created, Password, Image, [User Status].Name
		FROM User
		INNER JOIN [User Status] on User.Status = [User Status].ID`)
	checkErr(err)

	var users []User

	for rows.Next() {
		var user = User{}
		err = rows.Scan(&user.ID, &user.Name, &user.LastName, &user.UserName, &user.Email, &user.Created, &user.Password, &user.Image, &user.Status)
		checkErr(err)
		parsedDate, _ := time.Parse(dbDateLayout, user.Created)
		user.Created = parsedDate.Format(localDateLayout)
		user.Orders = getAllEntriesByUser(user, 3)
		users = append(users, user)
	}
	return users
}

func UserExists(name string) bool {
	user := User{}
	err := db.QueryRow(`
		SELECT ID
		FROM User
        WHERE [User Name] = $1`, user).Scan(&user.ID)

	return err == nil
}

func (user *User) Add() error {
	statement := "INSERT INTO User(Name, [Last Name], [User Name], Email, Created, Password, Image, Status) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"
	stmt, err := db.Prepare(statement)
	checkErr(err)

	defer stmt.Close()
	_, err = stmt.Exec("", "", user.UserName, user.Email, time.Now().Format(dbDateLayout), user.Password, "/img/user.svg", 1)
	return err
}

func (user *User) Update() error {
	var status int
	_ = db.QueryRow("SELECT ID FROM [User Status] WHERE Name = $1", user.Status).Scan(&status)

	_, err = db.Exec("UPDATE User SET [User Name]=$1, Email=$2, Password=$3, Status=$4 WHERE ID = $5", user.UserName, user.Email, user.Password, status, user.ID)
	return err
}

func (user *User) Delete() {
	_, err = db.Exec("DELETE FROM User WHERE ID = $1", user.ID)
	checkErr(err)
}

func GetStatus() []string {
	rows, err := db.Query("SELECT Name FROM [User Status]")
	checkErr(err)

	var status []string

	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		checkErr(err)
		status = append(status, s)
	}
	return status
}
