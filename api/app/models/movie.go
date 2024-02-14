package models

import "github.com/go-raptor/raptor"

type Movie struct {
	raptor.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
}

type Movies []Movie
