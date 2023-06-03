package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mtikvica/GoMovies/handlers"
)

func RegisterGenreRoutes(r *gin.Engine) {
	r.GET("/genres", handlers.GetGenres)
	r.POST("/genres", handlers.CreateGenre)
	r.PUT("/genres/:id", handlers.UpdateGenre)
	r.DELETE("/genres/:id", handlers.DeleteGenre)
}
