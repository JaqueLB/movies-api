package main

import (
	"moviesapi/controller"
	"moviesapi/router"
	"moviesapi/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	controller := &controller.MovieController{
		Storage: &storage.LocalStorage{},
	}
	engine := gin.Default()
	engine = router.CreateRoutes(engine, controller)
	engine.Run(":8080")
}
