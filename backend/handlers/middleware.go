package handlers

import (
	"backend/lib"
	"backend/models"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := c.Request().Header.Values("Authorization")
		token := strings.Split(strings.Join(data, ","), " ")[1]

		_, exist := models.CheckToken(lib.DB, token)
		if !exist {
			log.Println(exist)
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired jwt")
		}

		c.Response().Header().Set("x-token", token)

		return next(c)
	}
}
