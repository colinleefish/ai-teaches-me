package routes

import (
	handlers "gmdb/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.HandlePing)

	r.GET("/actors/", handlers.HandleGetActors)
	r.GET("/actors/:id", handlers.HandleGetActor)
	r.POST("/actors/", handlers.HandleCreateActor)
	r.PUT("/actors/:id", handlers.HandleUpdateActor)
	r.DELETE("/actors/:id", handlers.HandleDeleteActor)

	r.GET("/movies/", handlers.HandleGetMovies)
	r.GET("/movies/:id", handlers.HandleGetMovie)
	r.POST("/movies/", handlers.HandleCreateMovie)
	r.PUT("/movies/:id", handlers.HandleUpdateMovie)
	r.DELETE("/movies/:id", handlers.HandleDeleteMovie)

	r.GET("/awards/", handlers.HandleGetAwards)
	r.GET("/awards/:id", handlers.HandleGetAward)
	r.POST("/awards/", handlers.HandleCreateAward)
	r.PUT("/awards/:id", handlers.HandleUpdateAward)
	r.DELETE("/awards/:id", handlers.HandleDeleteAward)
}
