package language

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostLanguageHandler(context *gin.Context) {
	var newLanguage Language
	var err error

	if err = context.ShouldBindJSON(&newLanguage); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = ValidateLanguage(&newLanguage); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = CreateLanguage(&newLanguage); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newLanguage})
}

func GetLanguagesHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	languages, totalRecords, err := ReadAllLanguages(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": languages, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetLanguageHandler(context *gin.Context) {
	languageId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID format"})
		return
	}

	language, err := ReadOneLanguage(languageId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": language})
}

func PutLanguageHandler(context *gin.Context) {
	languageId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID format"})
		return
	}

	language, err := ReadOneLanguage(languageId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedLanguage Language
	err = context.ShouldBindJSON(&updatedLanguage)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language data format"})
		return
	}

	if err = ValidateLanguage(&updatedLanguage); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedLanguage.LanguageID = int(language.LanguageID)

	err = UpdateOneLanguage(updatedLanguage)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update language"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedLanguage})
}

func DeleteLanguageHandler(context *gin.Context) {
	languageId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID format"})
		return
	}

	language, err := ReadOneLanguage(languageId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = DeleteOneLanguage(*language)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete language"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": language})
}
