package lib

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id       int
	Username string
}

type JwtCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(id int, username string) (string, error) {

	claims := &JwtCustomClaims{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(SECRET_JWT))

	if err != nil {
		return "", err
	}

	return t, nil
}
