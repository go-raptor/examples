package services

import (
	"github.com/go-raptor/examples/api/app/models"
	"github.com/go-raptor/raptor"
)

type MoviesService struct {
	raptor.Service
}

func (ms *MoviesService) All() (models.Movies, error) {
	var movies models.Movies
	result := ms.DB.Find(&movies)
	return movies, result.Error
}

func (ms *MoviesService) Find(id int) (*models.Movie, error) {
	var movie models.Movie
	result := ms.DB.First(&movie, id)
	return &movie, result.Error
}
