package development

import (
	"main/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var countryObject = NewCountry(generateRandomString(51))
var cityObject = NewCity("1")

func GetTestHandler(context *gin.Context) {
	// makeTransaction(context)
}

func makeTransaction(context *gin.Context) {
	db.GORM.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&countryObject).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return err
		}

		cityObject.CountryID = int16(countryObject.CountryID)

		if err := tx.Create(&cityObject).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return err
		}

		return nil
	})

	context.JSON(http.StatusOK, gin.H{"message": "the test function has been executed"})
}
