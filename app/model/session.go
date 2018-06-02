package model

import (
	uuid "github.com/satori/go.uuid"
)

type Session struct {
	ID     string
	UserID uint32
}

func CreateSession(user User) Session {
	sID, err := uuid.NewV4()
	checkErr(err)

	session := Session{ID: sID.String(), UserID: user.ID}
	statement := "INSERT INTO Session(ID, [User ID]) VALUES($1, $2)"
	stmt, err := db.Prepare(statement)
	checkErr(err)

	defer stmt.Close()
	_, err = stmt.Exec(session.ID, user.ID)
	checkErr(err)

	return session
}

func GetSession(id string) (Session, error) {
	session := Session{}
	err := db.QueryRow(`
		SELECT Session.ID, Session.[User ID]
		FROM Session
		INNER JOIN User on Session.[User ID] = User.ID
        WHERE Session.ID = $1`, id).Scan(&session.ID, &session.UserID)

	return session, err
}

func DeleteSession(id string) {
	_, err = db.Exec("DELETE from Session where ID = $1", id)
	checkErr(err)
}

func SessionExists(id string) bool {
	session := Session{}
	err := db.QueryRow(`
		SELECT *
		FROM Session
		INNER JOIN User on Session.[User ID] = User.ID
        WHERE Session.ID = $1`, id).Scan(&session.ID, &session.UserID)

	return err == nil
}
