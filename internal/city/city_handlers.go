package city

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCityHandler(context *gin.Context) {
	var newCity City
	var err error

	if err = context.ShouldBindJSON(&newCity); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = ValidateCity(&newCity); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = CreateCity(db.GORM, &newCity); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newCity})
}

func GetCitiesHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	cities, totalRecords, err := ReadAllCities(db.GORM, pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": cities, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetCityHandler(context *gin.Context) {
	cityId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID format"})
		return
	}

	city, err := ReadOneCity(db.GORM, cityId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": city})
}

func PutCityHandler(context *gin.Context) {
	cityId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID format"})
		return
	}

	city, err := ReadOneCity(db.GORM, cityId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedCity City
	err = context.ShouldBindJSON(&updatedCity)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city data format"})
		return
	}

	if err = ValidateCity(&updatedCity); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCity.CityID = int(city.CityID)

	err = UpdateOneCity(db.GORM, updatedCity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update city"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedCity})
}

func DeleteCityHandler(context *gin.Context) {
	cityId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID format"})
		return
	}

	city, err := ReadOneCity(db.GORM, cityId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = DeleteOneCity(db.GORM, *city)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete city"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": city})
}
