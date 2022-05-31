package main

import (
	"moviesapi/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/movies", controllers.GetMovies)
	router.GET("/movies/:ID", controllers.GetMovieByID)
	router.PUT("/movies", controllers.CreateMovie)
	router.POST("/movies/:ID", controllers.UpdateMovieByID)
	router.DELETE("/movies/:ID", controllers.DeleteMovieByID)

	return router
}
