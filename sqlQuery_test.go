package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"golang-database/database"
	"testing"
	"time"
)

func TestQuerySql(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	queryselect := "SELECT id,name,email,balance,rating,birth_date,married,created_at FROM customer"

	rows, err := db.QueryContext(ctx, queryselect)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int32
		var balance int32
		var name string
		var email sql.NullString
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)

	}

	defer rows.Close()
}
