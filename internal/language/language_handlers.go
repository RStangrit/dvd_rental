package language

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LanguageHandler struct {
	service *LanguageService
}

func NewLanguageHandler(service *LanguageService) *LanguageHandler {
	return &LanguageHandler{service: service}
}

func (handler *LanguageHandler) PostLanguageHandler(context *gin.Context) {
	var newLanguage Language

	if err := context.ShouldBindJSON(&newLanguage); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.service.CreateLanguage(&newLanguage); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newLanguage})
}

func (handler *LanguageHandler) PostLanguagesHandler(context *gin.Context) {
	var newLanguages []*Language

	if err := context.ShouldBindJSON(&newLanguages); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErrors, createdLanguages, _ := handler.service.CreateLanguages(newLanguages)

	if len(validationErrors) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": validationErrors,
			"data":   createdLanguages,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": createdLanguages})
}

func (handler *LanguageHandler) GetLanguagesHandler(context *gin.Context) {
	var pagination db.Pagination

	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	//map filters, if no parameters are passed, the WHERE conditions are not included and the query will return all records
	filters := make(map[string]any)

	if name := context.Query("name"); name != "" {
		filters["name"] = name
	}

	languages, totalRecords, err := handler.service.ReadAllLanguages(pagination, filters)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": languages, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *LanguageHandler) GetLanguageHandler(context *gin.Context) {
	languageId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID format"})
		return
	}

	language, err := handler.service.ReadOneLanguage(languageId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": language})
}

func (handler *LanguageHandler) PutLanguageHandler(context *gin.Context) {
	languageId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID format"})
		return
	}

	var updatedLanguage Language
	if err := context.ShouldBindJSON(&updatedLanguage); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language data format"})
		return
	}

	updatedLanguage.LanguageID = int(languageId)

	if err := handler.service.UpdateOneLanguage(&updatedLanguage); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update language"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedLanguage})
}

func (handler *LanguageHandler) DeleteLanguageHandler(context *gin.Context) {
	languageId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID format"})
		return
	}

	language, err := handler.service.ReadOneLanguage(languageId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Language not found"})
		return
	}

	if err := handler.service.DeleteOneLanguage(language); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete language"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": language})
}
