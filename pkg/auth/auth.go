package auth

import (
	"fmt"
	"main/config"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = getSecretKey()
var tokenBlacklist = make(map[string]struct{})

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

	isTokenBlacklisted := isTokenBlacklisted(token.Raw)
	if isTokenBlacklisted {
		return fmt.Errorf("request a new token")
	}

	return nil
}

func GenerateHashFromPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func AddToBlacklist(token string) {
	tokenBlacklist[token] = struct{}{}
}

func isTokenBlacklisted(token string) bool {
	_, blacklisted := tokenBlacklist[token]
	return blacklisted
}
