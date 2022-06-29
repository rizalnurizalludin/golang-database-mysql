package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	myDB := "root:@tcp(localhost:3306)/golang_database?parseTime=true"
	db, err := sql.Open("mysql", myDB)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(60 * time.Second)

	return db
}
