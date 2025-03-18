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

// @Summary Create Language
// @Description Creates a new language
// @ID createLanguage
// @Tags Language
// @Accept json
// @Produce json
// @Param language body Language true "New language data"
// @Success 200 {object} Language "Created language"
// @Failure 400 {object} ErrorResponse "Validation error"
// @Router /language [post]
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

// @Summary Create Multiple Languages
// @Description Creates multiple languages
// @ID createLanguages
// @Tags Language
// @Accept json
// @Produce json
// @Param languages body []Language true "Array of languages"
// @Success 200 {object} LanguageListResponse "Created languages"
// @Failure 400 {object} ErrorResponse "Validation error"
// @Router /languages [post]
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

// @Summary Get All Languages
// @Description Retrieves a list of all languages with pagination
// @ID getLanguages
// @Tags Language
// @Produce json
// @Success 200 {object} LanguageListResponse "List of languages"
// @Router /languages [get]
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

// @Summary Get Language by ID
// @Description Retrieves a language by its identifier
// @ID getLanguageByID
// @Tags Language
// @Produce json
// @Param id path int true "Language ID"
// @Success 200 {object} Language "Found language"
// @Failure 404 {object} ErrorResponse "Language not found"
// @Router /language/{id} [get]
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

// @Summary Update Language
// @Description Updates a language by its ID
// @ID updateLanguage
// @Tags Language
// @Accept json
// @Produce json
// @Param id path int true "Language ID"
// @Param language body Language true "Data for update"
// @Success 200 {object} Language "Updated language"
// @Failure 400 {object} ErrorResponse "Validation error"
// @Router /language/{id} [put]
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

// @Summary Delete Language
// @Description Deletes a language by its ID
// @ID deleteLanguage
// @Tags Language
// @Produce json
// @Param id path int true "Language ID"
// @Success 200 {object} Language "Deleted language"
// @Failure 404 {object} ErrorResponse "Language not found"
// @Router /language/{id} [delete]
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
