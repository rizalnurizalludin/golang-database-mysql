package golangdatabase

import (
	"context"
	"fmt"
	"golang-database/database"
	"testing"
)

func TestSqlInjectionSafe(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin1"

	sqlQuery := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"

	fmt.Println(sqlQuery)

	rows, err := db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}
