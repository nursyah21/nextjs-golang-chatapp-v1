package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type Messages struct {
	Id         int    `json:"id"`
	SendId     string `json:"send_id"`
	ReceivedId string `json:"received_id"`
	Message    string `json:"message"`
	FileLink   string `json:"file_link"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

// type Message struct {
// 	Id         int
// 	SendId     int
// 	ReceivedId int
// 	Message
// }

// func CheckUser(db *sql.DB, username string) (res UsersCollection, exist bool) {
// 	sql := "SELECT id, username, password from users where username = ?"

// 	// Query for a value based on a single row.
// 	if err := db.QueryRow(sql, username).Scan(&res.Id, &res.Username, &res.Password); err != nil {
// 		log.Println(fmt.Printf("Error fetch data: %v", err.Error()))
// 		return res, false
// 	}

// 	return res, true
// }

// func CreateUser(db *sql.DB, username string, password string) bool {

// 	sql := `
// 	INSERT INTO
// 		users (username, password)
// 	SELECT ?, ?
// 	WHERE
// 		NOT EXISTS (
// 			SELECT 1
// 			FROM users
// 			WHERE
// 				username = ?
// 		);
// 	`

// 	// Create a prepared SQL statement
// 	stmt, err := db.Prepare(sql)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer stmt.Close()

// 	if _, err := stmt.Exec(username, password, username); err != nil {
// 		panic(err)
// 	}

// 	return false
// }
