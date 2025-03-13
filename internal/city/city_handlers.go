package city

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	service *CityService
}

func NewCityHandler(service *CityService) *CityHandler {
	return &CityHandler{service: service}
}

func (handler *CityHandler) PostCityHandler(context *gin.Context) {
	var newCity City
	var err error

	if err = context.ShouldBindJSON(&newCity); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = handler.service.CreateCity(&newCity); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newCity})
}

func (handler *CityHandler) GetCitiesHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	cities, totalRecords, err := handler.service.ReadAllCities(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": cities, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *CityHandler) GetCityHandler(context *gin.Context) {
	cityId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID format"})
		return
	}

	city, err := handler.service.ReadOneCity(cityId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": city})
}

func (handler *CityHandler) PutCityHandler(context *gin.Context) {
	cityId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID format"})
		return
	}

	city, err := handler.service.ReadOneCity(cityId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedCity *City
	err = context.ShouldBindJSON(&updatedCity)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city data format"})
		return
	}

	updatedCity.CityID = int(city.CityID)

	err = handler.service.UpdateOneCity(updatedCity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update city"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedCity})
}

func (handler *CityHandler) DeleteCityHandler(context *gin.Context) {
	cityId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID format"})
		return
	}

	city, err := handler.service.ReadOneCity(cityId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = handler.service.DeleteOneCity(city)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete city"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": city})
}
