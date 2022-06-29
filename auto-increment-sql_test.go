package golangdatabase

import (
	"context"
	"fmt"
	"golang-database/database"
	"testing"
)

func TestAutoIncrement(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "rizal@gmail.com"
	comment := "Test Komen"

	sqlQuery := "INSERT INTO comments(email,comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, sqlQuery, email, comment)
	if err != nil {
		panic(err)
	}
	InsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes insert new comment with id", InsertID)

}
