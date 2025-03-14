package film

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FilmHandler struct {
	service *FilmService
}

func NewFilmHandler(service *FilmService) *FilmHandler {
	return &FilmHandler{service: service}
}

func (handler *FilmHandler) PostFilmHandler(context *gin.Context) {
	var newFilm Film

	if err := context.ShouldBindJSON(&newFilm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.service.CreateFilm(&newFilm); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newFilm})
}

func (handler *FilmHandler) PostFilmsHandler(context *gin.Context) {
	var newFilms []*Film

	if err := context.ShouldBindJSON(&newFilms); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErrors, createdFilms, _ := handler.service.CreateFilms(newFilms)

	if len(validationErrors) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors, "data": createdFilms})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": createdFilms})
}

func (handler *FilmHandler) GetFilmsHandler(context *gin.Context) {
	var pagination db.Pagination
	var filters FilmFilter

	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	if err := context.ShouldBindQuery(&filters); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	films, totalRecords, err := handler.service.ReadAllFilms(pagination, filters)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": films, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *FilmHandler) GetFilmHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

	film, err := handler.service.ReadOneFilm(filmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": film})
}

func (handler *FilmHandler) GetFilmActorsHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

	filmActors, err := handler.service.ReadOneFilmActors(filmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": filmActors})
}

func (handler *FilmHandler) PutFilmHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

	var updatedFilm Film
	if err := context.ShouldBindJSON(&updatedFilm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film data format"})
		return
	}

	updatedFilm.FilmID = int(filmId)
	if err := handler.service.UpdateOneFilm(&updatedFilm); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update film"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedFilm})
}

func (handler *FilmHandler) DeleteFilmHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

	film, err := handler.service.ReadOneFilm(filmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := handler.service.DeleteOneFilm(film); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete film"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": film})
}

func (handler *FilmHandler) PostFilmDiscountHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

	film, _ := handler.service.ReadOneFilm(filmId)

	var body map[string]interface{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	discountValue, exists := body["discount"]
	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Discount field is required"})
		return
	}

	discount, ok := discountValue.(float64)
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Discount must be a number"})
		return
	}

	if err := handler.service.DiscountOneFilm(filmId, discount); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set film discount: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": film, "discount": discount})
}
