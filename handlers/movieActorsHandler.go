package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtikvica/GoMovies/db"
	"github.com/mtikvica/GoMovies/models"
)

func CreateMovieActor(c *gin.Context) {
	// Get the movie ID from the URL parameter
	movieID := c.Param("id")

	// Parse the actor ID from the request body
	var requestBody struct {
		ActorID int `json:"actor_id"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	_, err := db.DB.Exec("INSERT INTO movie_actors (movie_id, actor_id) VALUES ($1, $2)", movieID, requestBody.ActorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie actor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie actor created successfully"})
}

func GetMovieActors(c *gin.Context) {
	movieID := c.Param("id")

	rows, err := db.DB.Query("SELECT actors.id, actors.name FROM actors INNER JOIN movie_actors ON actors.id = movie_actors.actor_id WHERE movie_actors.movie_id = $1", movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movie actors"})
		return
	}
	defer rows.Close()

	movieActors := []models.Actor{}

	for rows.Next() {
		var actor models.Actor
		err := rows.Scan(&actor.ID, &actor.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movie actors"})
			return
		}
		movieActors = append(movieActors, actor)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movie actors"})
		return
	}

	c.JSON(http.StatusOK, movieActors)
}
