package golangdatabase

import (
	"context"
	"fmt"
	"golang-database/database"
	"strconv"
	"testing"
)

func TestPrepareStatemen(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := "INSERT INTO comments(email,comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer statement.Close()
	for i := 0; i < 10; i++ {
		email := "rizal" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment ke :" + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		LastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id", LastInsertId)
	}
}

func TestTransaction(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	//transaction
	sqlQuery := "INSERT INTO comments(email,comment) VALUES(?,?)"
	for i := 0; i < 10; i++ {
		email := "rizal" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment ke :" + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, sqlQuery, email, comment)
		if err != nil {
			panic(err)
		}

		LastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id", LastInsertId)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
