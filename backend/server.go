package main

import (
	"backend/handlers"
	"backend/lib"
	"net/http"

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

	db := lib.InitDb()

	e.POST("/login", handlers.Login(db))
	e.POST("/register", handlers.Register(db))

	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := echojwt.Config{

		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(lib.JwtCustomClaims)
		},
		SigningKey: []byte(lib.SECRET_JWT),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("", restricted)

	// e.GET("/tasks", handlers.GetTasks(db))
	// e.PUT("/tasks", handlers.PutTask(db))
	// e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// e.GET("/users/:id", getUser)
	// e.GET("/show", show)
	// e.POST("/save", save)

	e.Logger.Fatal(e.Start(":5000"))
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*lib.JwtCustomClaims)
	username := claims.Username
	return c.String(http.StatusOK, "Welcome "+username+"!")
}
