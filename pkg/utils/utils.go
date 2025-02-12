package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetIntParam(context *gin.Context, intParamName string) (int64, error) {
	paramValue := context.Param(intParamName)
	if paramValue == "" {
		return 0, fmt.Errorf("parameter %s is empty", intParamName)
	}
	intValue, err := strconv.ParseInt(paramValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid integer value for parameter %s: %s", intParamName, paramValue)
	}
	return intValue, nil
}

func JoinStrings(strs ...string) string {
	return strings.Join(strs, " ")
}
