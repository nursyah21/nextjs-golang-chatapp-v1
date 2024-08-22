package handlers

import (
	"backend/models"
	"database/sql"
	"encoding/base64"
	"net/http"

	"crypto/rand"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/argon2"
)

func Login(c echo.Context) error {
	return c.String(http.StatusOK, "login")
}

func Register(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.Users
		c.Bind(&user)

		encPassword, err := generatePassword(user.Password)
		if err != nil {
			return c.String(http.StatusBadRequest, "error create account")
		}

		_, err2 := models.CreateUser(db, user.Username, string(encPassword))

		if err2 != nil {
			return c.String(http.StatusBadRequest, "error create account")
		}

		return c.String(http.StatusOK, "create account success")
	}
}

func generatePassword(password string) (hash string, err error) {

	salt, err := generateRandomBytes(16)
	if err != nil {
		return "", err
	}

	_hash := argon2.IDKey([]byte(password), salt, 3, 32*1024, 4, 32)
	hash = base64.RawStdEncoding.EncodeToString([]byte(_hash))

	return hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
