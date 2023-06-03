package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mtikvica/GoMovies/handlers"
)

func RegisterActorRoutes(r *gin.Engine) {
	r.GET("/actors", handlers.GetActors)
	r.GET("/actors/:id", handlers.GetActor)
	r.POST("/actors", handlers.CreateActor)
	r.PUT("/actors/:id", handlers.UpdateActor)
	r.DELETE("/actors/:id", handlers.DeleteActor)
}
