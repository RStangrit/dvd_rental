package handlers

import (
	"main/internal/models"
	"main/internal/repositories"
	"main/internal/services"
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCountryHandler(context *gin.Context) {
	var newCountry models.Country
	var err error

	if err = context.ShouldBindJSON(&newCountry); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = services.ValidateCountry(&newCountry); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = repositories.CreateCountry(&newCountry); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newCountry})
}

func GetCountriesHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	countries, totalRecords, err := repositories.ReadAllCountries(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": countries, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetCountryHandler(context *gin.Context) {
	countryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID format"})
		return
	}

	country, err := repositories.ReadOneCountry(countryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": country})
}

func PutCountryHandler(context *gin.Context) {
	countryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID format"})
		return
	}

	country, err := repositories.ReadOneCountry(countryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedCountry models.Country
	err = context.ShouldBindJSON(&updatedCountry)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country data format"})
		return
	}

	if err = services.ValidateCountry(&updatedCountry); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCountry.CountryID = country.CountryID

	err = repositories.UpdateOneCountry(updatedCountry)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update country"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedCountry})
}

func DeleteCountryHandler(context *gin.Context) {
	countryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID format"})
		return
	}

	country, err := repositories.ReadOneCountry(countryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = repositories.DeleteOneCountry(*country)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete country"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": country})
}
