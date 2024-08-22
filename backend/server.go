package main

import (
	"backend/handlers"
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := initDb("database.db")
	migrateDb(db)

	e.POST("/login", handlers.Login)
	e.POST("/register", handlers.Register(db))

	// e.GET("/tasks", handlers.GetTasks(db))
	// e.PUT("/tasks", handlers.PutTask(db))
	// e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// e.GET("/users/:id", getUser)
	// e.GET("/show", show)
	// e.POST("/save", save)

	e.Logger.Fatal(e.Start(":5000"))
}

// func register(c echo.Context) error {
// 	// Get name and email
// 	name := c.FormValue("username")
// 	email := c.FormValue("password")
// 	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
// }

func initDb(filePath string) *sql.DB {

	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrateDb(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(60) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		avatar VARCHAR(255),
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	-- Create the messages table
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		send_id INTEGER NOT NULL REFERENCES users (id),
		receive_id INTEGER NOT NULL REFERENCES users (id),
		message VARCHAR(512) NOT NULL,
		file_link VARCHAR(255),
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	-- Create the refresh_tokens table
	CREATE TABLE IF NOT EXISTS refresh_tokens (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL REFERENCES users (id),
		token VARCHAR(255) NOT NULL,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TRIGGER IF NOT EXISTS update_users_updated_at
	AFTER UPDATE ON users
	WHEN old.updatedAt <> current_timestamp
	BEGIN
		UPDATE users
		SET updatedAt = CURRENT_TIMESTAMP
		WHERE id = OLD.id;
	END;

	CREATE TRIGGER IF NOT EXISTS update_messages_updated_at
	AFTER UPDATE ON messages
	WHEN old.updatedAt <> current_timestamp
	BEGIN
		UPDATE messages
		SET updatedAt = CURRENT_TIMESTAMP
		WHERE id = OLD.id;
	END;

	CREATE TRIGGER IF NOT EXISTS update_refresh_tokens_updated_at
	AFTER UPDATE ON refresh_tokens
	WHEN old.updatedAt <> current_timestamp
	BEGIN
		UPDATE refresh_tokens
		SET updatedAt = CURRENT_TIMESTAMP
		WHERE id = OLD.id;
	END;
	`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
