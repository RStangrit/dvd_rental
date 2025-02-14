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

	if err = CreateLanguage(db.GORM, &newLanguage); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newLanguage})
}

func PostLanguagesHandler(context *gin.Context) {
	var newLanguages []Language
	var err error

	if err = context.ShouldBindJSON(&newLanguages); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var validationErrors []string
	var createdLanguages []Language

	for _, newLanguage := range newLanguages {
		if err = ValidateLanguage(&newLanguage); err != nil {
			validationErrors = append(validationErrors, err.Error())
			continue
		}

		if err = CreateLanguage(db.GORM, &newLanguage); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		createdLanguages = append(createdLanguages, newLanguage)
	}

	if len(validationErrors) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": validationErrors,
			"data":   createdLanguages,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": createdLanguages})
}

func GetLanguagesHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	//map filters, if no parameters are passed, the WHERE conditions are not included and the query will return all records
	filters := make(map[string]any)

	if name := context.Query("name"); name != "" {
		filters["name"] = name
	}

	languages, totalRecords, err := ReadAllLanguages(db.GORM, pagination, filters)
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

	language, err := ReadOneLanguage(db.GORM, languageId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": language})
}

func GetLanguageAssociatedFilmsHandler(context *gin.Context) {
	languageId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID format"})
		return
	}

	language, err := ReadOneLanguage(db.GORM, languageId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = language.LoadFilms(db.GORM)
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

	language, err := ReadOneLanguage(db.GORM, languageId)
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

	err = UpdateOneLanguage(db.GORM, updatedLanguage)
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

	language, err := ReadOneLanguage(db.GORM, languageId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = DeleteOneLanguage(db.GORM, *language)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete language"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": language})
}
