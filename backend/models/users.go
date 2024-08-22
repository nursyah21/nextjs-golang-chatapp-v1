package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UsersCollection struct {
	Username string
	Password string
}

func CheckUser(db *sql.DB, username string) (res UsersCollection, error bool) {
	sql := "SELECT username, password from users where username = ?"

	// Query for a value based on a single row.
	if err := db.QueryRow(sql, username).Scan(&res.Username, &res.Password); err != nil {
		log.Println(fmt.Printf("Error fetch data: %v", err.Error()))
		return res, false
	}
	fmt.Println(res)
	return res, true
}

func CreateUser(db *sql.DB, username string, password string) bool {

	sql := `
	INSERT INTO
		users (username, password)
	SELECT ?, ?
	WHERE
		NOT EXISTS (
			SELECT 1
			FROM users
			WHERE
				username = ?
		);
	`

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(username, password, username); err != nil {
		panic(err)
	}

	return false
}
