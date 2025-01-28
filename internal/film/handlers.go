package film

import (
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
