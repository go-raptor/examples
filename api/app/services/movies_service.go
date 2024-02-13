package services

import (
	"github.com/go-raptor/examples/api/app/models"
	"github.com/go-raptor/raptor"
	"golang.org/x/exp/maps"
)

type MoviesService struct {
	raptor.Service

	movies map[int]models.Movie
}

func NewMoviesService() *MoviesService {
	return &MoviesService{
		movies: map[int]models.Movie{
			1: {
				Title:       "The Shawshank Redemption",
				Description: "A gripping tale of hope and redemption as a man finds solace and purpose in the confines of a prison.",
				Year:        1994,
			},
			2: {
				Title:       "The Godfather",
				Description: "A powerful crime family saga that follows the transformation of Michael Corleone from a reluctant family outsider to a ruthless mafia boss.",
				Year:        1972,
			},
		},
	}
}

func (ms *MoviesService) All() []models.Movie {
	return maps.Values(ms.movies)
}

func (ms *MoviesService) Find(id int) *models.Movie {
	if movie, ok := ms.movies[id]; ok {
		return &movie
	}

	return nil
}
