package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mtikvica/GoMovies/handlers"
)

func RegisterMoviesActorsRoutes(r *gin.Engine) {
	r.GET("/movies/:id/actors", handlers.GetMovieActors)
	r.POST("/movies/:id/actors", handlers.CreateMovieActor)
}
