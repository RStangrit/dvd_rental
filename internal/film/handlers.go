package film

import (
	"main/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postFilmHandler(context *gin.Context) {
	var newFilm Film
	var err error

	if err = context.ShouldBindJSON(&newFilm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = newFilm.Validate(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = newFilm.createFilm(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	context.JSON(http.StatusCreated, gin.H{"data": newFilm})
}

func getFilmshandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	films, totalRecords, err := readAllFilms(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": films, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}
