package initializers

import (
	"github.com/go-raptor/example-app/app/controllers"
	"github.com/go-raptor/example-app/app/middlewares"
	"github.com/go-raptor/example-app/app/services"
	"github.com/go-raptor/raptor"
)

func App() *raptor.AppInitializer {
	ms := services.NewMoviesService()

	return &raptor.AppInitializer{
		Services: raptor.Services{
			ms,
		},

		Middlewares: raptor.Middlewares{
			&middlewares.ActionLogger{},
		},

		Controllers: raptor.Controllers{
			&controllers.HomeController{},
			&controllers.MoviesController{
				Ms: ms,
			},
		},
	}
}
