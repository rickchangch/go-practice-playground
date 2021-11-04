package suitepkg

import (
	"database/sql"
	"go-practice-playground/localTest/db"
)

func Get(dbc *sql.DB, id string) (*db.User, error) {
	user := &db.User{}
	err := dbc.QueryRow("SELECT * FROM user WHERE UserID = ?", id).Scan(&user.UserID, &user.Name, &user.Created)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Add(dbc *sql.DB, user db.User) error {
	stmt, err := dbc.Prepare("INSERT INTO user(UserID, Name, Created) values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.UserID, user.Name, user.Created)
	if err != nil {
		return err
	}

	return nil
}
