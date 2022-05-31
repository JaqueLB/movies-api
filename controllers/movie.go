package controllers

import (
	"moviesapi/entities"
	"moviesapi/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovies(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": storage.List(),
	})
}

func GetMovieByID(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	movie := storage.GetByID(ID)

	ctx.JSON(http.StatusOK, gin.H{
		"data": movie,
	})
}

func UpdateMovieByID(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("ID"))
	var movie *entities.Movie

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	storage.UpdateByID(ID, movie)
	ctx.JSON(http.StatusOK, gin.H{
		"data": "OK",
	})
}

func DeleteMovieByID(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	storage.DeleteByID(ID)

	ctx.JSON(http.StatusOK, gin.H{
		"data": "OK",
	})
}

func CreateMovie(ctx *gin.Context) {
	var movie *entities.Movie

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	movie = storage.Create(movie)

	ctx.JSON(http.StatusOK, gin.H{
		"data": movie,
	})
}
