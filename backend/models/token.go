package models

import (
	"database/sql"
	"log"

	// "fmt"
	// "log"

	_ "github.com/mattn/go-sqlite3"
)

type Token struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Token     string `json:"token"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type TokenCollection struct {
	token string
}

func CheckToken(db *sql.DB, token string) (res TokenCollection, exist bool) {
	sql := "SELECT 1 from refresh_tokens where token = ?"

	// Query for a value based on a single row.
	if err := db.QueryRow(sql, token).Scan(&res.token); err != nil {
		log.Printf("Error fetch data: %v", err.Error())
		return res, false
	}

	return res, true
}

func DeleteToken(db *sql.DB, token string) bool {

	sql := `
		DELETE FROM refresh_tokens WHERE token = ?;
	`

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	if _, err := stmt.Exec(token); err != nil {
		panic(err.Error())
	}

	return false
}

func StoreToken(db *sql.DB, user_id int, token string) bool {

	sql := `
	INSERT INTO
		refresh_tokens (user_id, token)
	VALUES (?, ?)
	`

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	if _, err := stmt.Exec(user_id, token); err != nil {
		panic(err.Error())
	}

	return false
}
