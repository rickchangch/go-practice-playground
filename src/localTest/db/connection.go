package db

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	USER         string        = "root"
	PWD          string        = "rootpass"
	ADDRESS      string        = "localhost"
	PORT         int           = 3306
	DB           string        = "mytest"
	PARAM        string        = "charset=utf8mb4,utf8&parseTime=true"
	MaxLifetime  time.Duration = 10
	MaxOpenConns int           = 10
	MaxIdleConns int           = 5
)

func Connect() (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", USER, PWD, ADDRESS, PORT, DB, PARAM)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(MaxLifetime * time.Second)
	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)

	return db, nil
}
