package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mtikvica/GoMovies/handlers"
)

func RegisterMovieRoutes(r *gin.Engine) {
	r.GET("/movies", handlers.GetMovies)
	r.GET("/movies/:id", handlers.GetMovie)
	r.GET("/movies/genre/:genre", handlers.GetMoviesByGenre)
	r.POST("/movies", handlers.CreateMovie)
	r.PUT("/movies/:id", handlers.UpdateMovie)
	r.DELETE("/movies/:id", handlers.DeleteMovie)
}
