package models

import "time"

type Movie struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
	GenreID  int    `json:"genre_id"`
	Genre    Genre
}

type Actor struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
}

type MovieActor struct {
	MovieID int `json:"movie_id"`
	Movie   Movie
	ActorID int `json:"actor_id"`
	Actor   Actor
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
