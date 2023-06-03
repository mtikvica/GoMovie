package main

import (
	"github.com/mtikvica/GoMovies/handlers"
	"github.com/mtikvica/GoMovies/routes"

	"github.com/mtikvica/GoMovies/db"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := db.ConnectDB(); err != nil {
		panic(err)
	}

	r := gin.Default()

	routes.RegisterMovieRoutes(r)

	routes.RegisterActorRoutes(r)

	r.GET("/movies/:id/actors", handlers.GetMovieActors)
	r.POST("/movies/:id/actors", handlers.CreateMovieActor)

	r.Run(":8000")
}
