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

	// db := lib.InitDb()

	e.POST("/login", handlers.Login())
	e.POST("/register", handlers.Register())
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(lib.JwtCustomClaims)
		},
		SigningKey: []byte(lib.SECRET_JWT),
	}

	r.Use(echojwt.WithConfig(config))
	r.Use(handlers.Middleware)

	// r.GET("", restricted)
	r.DELETE("/logout", handlers.Logout)

	e.Logger.Fatal(e.Start(":5000"))
}
