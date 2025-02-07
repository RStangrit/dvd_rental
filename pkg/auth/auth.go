package auth

import (
	"fmt"
	"main/config"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = getSecretKey()

func getSecretKey() []byte {
	params := config.LoadConfig()
	secretKey := []byte(params.SecretKey)
	return secretKey
}

func CreateToken(email string) (string, error) {
	var secretKey = getSecretKey()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 2).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
