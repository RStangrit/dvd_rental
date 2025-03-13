package country

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	service *CountryService
}

func NewCountryHandler(service *CountryService) *CountryHandler {
	return &CountryHandler{service: service}
}

func (handler *CountryHandler) PostCountryHandler(context *gin.Context) {
	var newCountry Country
	var err error

	if err = context.ShouldBindJSON(&newCountry); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = handler.service.CreateCountry(&newCountry); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newCountry})
}

func (handler *CountryHandler) GetCountriesHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	countries, totalRecords, err := handler.service.ReadAllCountries(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": countries, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *CountryHandler) GetCountryHandler(context *gin.Context) {
	countryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID format"})
		return
	}

	country, err := handler.service.ReadOneCountry(countryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": country})
}

func (handler *CountryHandler) PutCountryHandler(context *gin.Context) {
	countryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID format"})
		return
	}

	country, err := handler.service.ReadOneCountry(countryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedCountry *Country
	err = context.ShouldBindJSON(&updatedCountry)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country data format"})
		return
	}

	updatedCountry.CountryID = country.CountryID

	err = handler.service.UpdateOneCountry(updatedCountry)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update country"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedCountry})
}

func (handler *CountryHandler) DeleteCountryHandler(context *gin.Context) {
	countryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID format"})
		return
	}

	country, err := handler.service.ReadOneCountry(countryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = handler.service.DeleteOneCountry(country)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete country"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": country})
}
