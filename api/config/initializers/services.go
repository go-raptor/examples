package initializers

import (
	"github.com/go-raptor/examples/api/app/services"
	"github.com/go-raptor/raptor"
)

func Services() raptor.Services {
	return raptor.Services{
		services.NewMoviesService(),
	}
}
