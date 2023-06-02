package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mtikvica/GoMovies/db"
	"github.com/mtikvica/GoMovies/models"
)

func GetMovies(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, title, director, year FROM movies")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movies"})
		return
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Director, &movie.Year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movies"})
			return
		}
		movies = append(movies, movie)
	}

	c.JSON(http.StatusOK, movies)
}

func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(movie)

	_, err := db.DB.Exec("INSERT INTO movies (title, director, year) VALUES ($1, $2, $3)",
		movie.Title, movie.Director, movie.Year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func GetMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var movie models.Movie
	err = db.DB.QueryRow("SELECT id, title, director, year FROM movies WHERE id = $1", id).Scan(
		&movie.ID, &movie.Title, &movie.Director, &movie.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movie"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func UpdateMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var updatedMovie models.Movie
	if err := c.ShouldBindJSON(&updatedMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.DB.Exec("UPDATE movies SET title = $1, director = $2, year = $3 WHERE id = $4",
		updatedMovie.Title, updatedMovie.Director, updatedMovie.Year, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.Status(http.StatusOK)
}

func DeleteMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	result, err := db.DB.Exec("DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.Status(http.StatusOK)
}
