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

	Ms *services.MoviesService
}

func (mc *MoviesController) Index(c *raptor.Context) error {
	ms := mc.Services["MoviesService"].(*services.MoviesService)

	return c.JSON(ms.All())
}

func (mc *MoviesController) Get(c *raptor.Context) error {
	ms := mc.Services["MoviesService"].(*services.MoviesService)
	id, err := c.ParamsInt("id")
	if err == nil {
		if movie := ms.Find(id); movie != nil {
			return c.JSON(movie)
		}
	}

	return c.JSON(models.NewError(http.StatusNotFound, fmt.Sprintf("Movie with id %d not found", id)), http.StatusNotFound)
}
