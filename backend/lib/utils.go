package lib

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func GetId(c echo.Context) (any, error) {
	data := c.Response().Header().Values("x-token")
	token := strings.Join(data, ",")

	res, err := ParseToken(token)
	if err != nil {
		return "", err
	}

	return res, nil
}
