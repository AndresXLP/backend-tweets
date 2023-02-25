package utils

import (
	"fmt"
	"time"

	"github.com/andresxlp/backend-twitter/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
)

type JWT interface {
	GenerateToken(data string) (string, error)
	ValidateToken(receivedToken string) (string, error)
}

type jwtUtils struct{}

func NewJWTUtils() JWT {
	return &jwtUtils{}
}

type customClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var mySecret = []byte(config.Environments().SecretJWT)

func (u *jwtUtils) GenerateToken(data string) (string, error) {
	claim := customClaim{
		Email: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(mySecret)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	return tokenString, nil
}

func (u *jwtUtils) ValidateToken(receivedToken string) (string, error) {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return mySecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["email"].(string), nil
	}

	return "", fmt.Errorf("token invalid")
}
