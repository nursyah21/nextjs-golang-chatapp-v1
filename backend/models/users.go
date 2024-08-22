package models

import (
	"database/sql"

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

// func UserExist(db *sql.DB, username string) bool {
// 	sql := "SELECT username FROM users  WHERE username = ?"

// 	// Create a prepared SQL statement
// 	stmt, err := db.Prepare(sql)
// 	// Exit if we get an error
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Make sure to cleanup after the program exits
// 	defer stmt.Close()

// 	// Replace the '?' in our prepared statement with 'name'
// 	result, err2 := stmt.Exec(username)
// 	// Exit if we get an error
// 	if err2 != nil {
// 		panic(err2)
// 	}
// 	exist, _ := result.LastInsertId()
// 	if exist != 0 {
// 		return true
// 	}

// 	return false
// }

func CreateUser(db *sql.DB, username string, password string) (int64, error) {
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

	result, err2 := stmt.Exec(username, password, username)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}
