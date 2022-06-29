package golangdatabase

import (
	"context"
	"fmt"
	"golang-database/database"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestExecSql(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	data := "INSERT INTO customer(name,email,balance,rating,birth_date) VALUES('Nurizalludin','nurizalludin@gmail.com',1000000000,9.9,'1996-03-17')"

	_, err := db.ExecContext(ctx, data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes insert new Customer")
}

func TestExecSqlParameter(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()
	username := "rizal;DROP FROM user;#"
	password := "rizal"

	ctx := context.Background()
	data := "INSERT INTO user(username,password) VALUES(?,?)"

	_, err := db.ExecContext(ctx, data, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes insert new Customer")
}
