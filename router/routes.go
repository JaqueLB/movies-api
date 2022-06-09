package router

import (
	"moviesapi/controller"

	"github.com/gin-gonic/gin"
)

func CreateRoutes(router *gin.Engine, controller *controller.MovieController) *gin.Engine {
	router.GET("/movies", controller.GetMovies)
	router.GET("/movies/:ID", controller.GetMovieByID)
	router.PUT("/movies", controller.CreateMovie)
	router.POST("/movies/:ID", controller.UpdateMovieByID)
	router.DELETE("/movies/:ID", controller.DeleteMovieByID)

	return router
}
