package main

import (
	"backend/handlers"
	"backend/lib"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	r := e.Group("/restricted")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(lib.JwtCustomClaims)
		},
		SigningKey: []byte(lib.SECRET_JWT),
	}
	r.Use(echojwt.WithConfig(config))
	r.Use(handlers.Middleware)

	w := r.Group("/ws")

	e.POST("/login", handlers.Login())
	e.POST("/register", handlers.Register())
	r.DELETE("/logout", handlers.Logout)

	r.GET("/message", handlers.ReadMessage)
	r.POST("/message", handlers.SendMessage)
	r.PUT("/message", handlers.EditMessage)
	r.DELETE("/message", handlers.DeleteFile)

	r.PUT("/profile", handlers.EditProfile)
	r.GET("/profile", handlers.ReadProfile)

	r.GET("/files", handlers.ReadFile)
	r.POST("/files", handlers.SendFile)
	r.DELETE("/files", handlers.DeleteFile)

	r.GET("/user/:id", handlers.SearchUser)

	w.GET("/message", handlers.WebSocketMessage)

	e.Logger.Fatal(e.Start(":5000"))
}
