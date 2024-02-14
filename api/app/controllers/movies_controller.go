package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-raptor/examples/api/app/models"
	"github.com/go-raptor/examples/api/app/services"
	"github.com/go-raptor/raptor"
)

type MoviesController struct {
	raptor.Controller
}

func (mc *MoviesController) Movies() *services.MoviesService {
	return mc.Services["MoviesService"].(*services.MoviesService)
}

func (mc *MoviesController) Index(c *raptor.Context) error {
	movies, err := mc.Movies().All()
	if err != nil {
		return c.JSON(models.NewError(http.StatusInternalServerError, err.Error()), http.StatusInternalServerError)
	}
	return c.JSON(movies)
}

func (mc *MoviesController) Get(c *raptor.Context) error {
	ms := mc.Services["MoviesService"].(*services.MoviesService)
	id, err := c.ParamsInt("id")
	if err == nil {
		if movie, _ := ms.Find(id); movie != nil {
			return c.JSON(movie)
		}
	}

	return c.JSON(models.NewError(http.StatusNotFound, fmt.Sprintf("Movie with id %d not found", id)), http.StatusNotFound)
}
