package auth

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var AppKey = os.Getenv("APP_KEY")

func CreateToken(userID int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":     time.Now().Unix()})

	tokenString, err := token.SignedString([]byte(AppKey))

	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	return tokenString, nil
}

func RefreshToken(userID int) (interface{}, error) {
	return nil, nil
}
