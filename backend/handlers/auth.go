package handlers

import (
	"backend/lib"
	"backend/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.Users
		c.Bind(&user)

		res, exist := models.CheckUser(db, user.Username)
		if !exist {
			return c.String(http.StatusOK, fmt.Sprintf("user: %v  not exist", res.Username))
		}

		match, _ := lib.ComparePasswordAndHash(user.Password, res.Password)
		if !match {
			return c.String(http.StatusBadRequest, "wrong password")
		}

		t, _ := lib.CreateToken(res.Id, res.Username)

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})

		// if err2 != nil {
		// 	return c.String(http.StatusBadRequest, "error create account")
		// }

		// var i models.UsersCollection
		// return c.String(http.StatusOK, fmt.Sprintf("user: %v, password: %v", res.Username, res.Password))
	}
}

func Register(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.Users
		c.Bind(&user)

		p := &lib.Params{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		}

		encPassword, err := lib.GeneratePassword(user.Password, p)
		if err != nil {
			return c.String(http.StatusBadRequest, "error create account")
		}
		if _, exist := models.CheckUser(db, user.Username); exist {
			return c.String(http.StatusBadRequest, "account already exist")
		}
		models.CreateUser(db, user.Username, string(encPassword))
		return c.String(http.StatusOK, "create account success")

	}
}
