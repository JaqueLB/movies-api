package controller

import (
	"moviesapi/external"
	"moviesapi/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	Storage storage.IStorage
}

func (c *MovieController) GetMovies(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": c.Storage.List(),
	})
}

func (c *MovieController) GetMovieByID(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	movie := c.Storage.GetByID(ID)

	ctx.JSON(http.StatusOK, gin.H{
		"data": movie,
	})
}

func (c *MovieController) UpdateMovieByID(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("ID"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	var data *external.MovieRequest

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.Storage.UpdateByID(ID, data)
	ctx.JSON(http.StatusOK, gin.H{
		"data": "OK",
	})
}

func (c *MovieController) DeleteMovieByID(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.Storage.DeleteByID(ID)

	ctx.JSON(http.StatusOK, gin.H{
		"data": "OK",
	})
}

func (c *MovieController) CreateMovie(ctx *gin.Context) {
	var data *external.MovieRequest

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	movie := c.Storage.Create(data)

	ctx.JSON(http.StatusOK, gin.H{
		"data": movie,
	})
}
