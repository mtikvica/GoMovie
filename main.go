package main

import (
	"github.com/mtikvica/GoMovies/handlers"

	"github.com/mtikvica/GoMovies/db"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := db.ConnectDB(); err != nil {
		panic(err)
	}

	r := gin.Default()

	// Movie endpoints
	r.GET("/movies", handlers.GetMovies)
	r.POST("/movies", handlers.CreateMovie)
	r.GET("/movies/:id", handlers.GetMovie)
	r.PUT("/movies/:id", handlers.UpdateMovie)
	r.DELETE("/movies/:id", handlers.DeleteMovie)

	r.Run(":8000")
}
