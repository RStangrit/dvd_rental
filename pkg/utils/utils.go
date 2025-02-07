package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetIntParam(context *gin.Context, intParamName string) (int64, error) {
	intValue, err := strconv.ParseInt(context.Param(intParamName), 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(intValue), nil
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
