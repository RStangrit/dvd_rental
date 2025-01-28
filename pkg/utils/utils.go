package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIntParam(context *gin.Context, intParamName string) (int64, error) {
	intValue, err := strconv.ParseInt(context.Param(intParamName), 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(intValue), nil
}
