package film

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostFilmHandler(context *gin.Context) {
	var newFilm Film
	var err error

	if err = context.ShouldBindJSON(&newFilm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = ValidateFilm(&newFilm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = CreateFilm(&newFilm); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newFilm})
}

func PostFilmsHandler(context *gin.Context) {
	var newFilms []Film
	var err error

	if err = context.ShouldBindJSON(&newFilms); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var validationErrors []string
	var createdFilms []Film

	for _, newFilm := range newFilms {
		if err = ValidateFilm(&newFilm); err != nil {
			validationErrors = append(validationErrors, err.Error())
			continue
		}

		if err = CreateFilm(&newFilm); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		createdFilms = append(createdFilms, newFilm)
	}

	if len(validationErrors) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": validationErrors,
			"data":   createdFilms,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": createdFilms})
}

func GetFilmshandler(context *gin.Context) {
	var pagination db.Pagination
	var err error
	//struct filters, makes code more clear and works only with non-empty values
	var filters FilmFilter

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	if err = context.ShouldBindQuery(&filters); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	films, totalRecords, err := ReadAllFilms(pagination, filters)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": films, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetFilmHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

	film, err := ReadOneFilm(filmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": film})
}

func GetFilmActorsHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

	filmWIthActors, err := ReadOneFilmActors(filmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": filmWIthActors})
}

func PutFilmHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

	film, err := ReadOneFilm(filmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedFilm Film
	err = context.ShouldBindJSON(&updatedFilm)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film data format"})
		return
	}

	if err = ValidateFilm(&updatedFilm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedFilm.FilmID = int(film.FilmID)

	err = UpdateOneFilm(updatedFilm)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update film"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedFilm})
}

func DeleteFilmHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

	film, err := ReadOneFilm(filmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = DeleteOneFilm(*film)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete film"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": film})
}

func PostFilmDiscountHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

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

	if err := ValidateDiscountPercentage(discount); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	film, err := ReadOneFilm(filmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if film == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Film not found"})
		return
	}

	if err := DiscountOneFilm(*film, discount); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set film discount " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": film, "discount": discount})
}
