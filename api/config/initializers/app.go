package initializers

import (
	"github.com/go-raptor/examples/api/db"
	"github.com/go-raptor/raptor"
)

/*func App() *raptor.AppInitializer {
	ms := services.NewMoviesService()

	return &raptor.AppInitializer{
		Database: raptor.Migrations{
			1: migrations.AddMoviesTable,
		},

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
}*/

func App() *raptor.AppInitializer {
	return &raptor.AppInitializer{
		Database:    db.Migrations(),
		Services:    Services(),
		Middlewares: Middlewares(),
		Controllers: Controllers(),
	}
}
