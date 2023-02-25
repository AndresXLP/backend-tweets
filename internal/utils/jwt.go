package utils

import (
	"time"

	"github.com/andresxlp/backend-twitter/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
)

type customClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(data string) (string, error) {
	claim := customClaim{
		Email: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	mySecret := []byte(config.Environments().SecretJWT)
	tokenString, err := token.SignedString(mySecret)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	return tokenString, nil
}
