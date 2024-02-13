package initializers

import (
	"github.com/go-raptor/examples/api/app/controllers"
	"github.com/go-raptor/raptor"
)

func Controllers() raptor.Controllers {
	return raptor.Controllers{
		&controllers.HomeController{},
		&controllers.MoviesController{},
	}
}
