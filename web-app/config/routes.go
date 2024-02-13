package config

import "github.com/go-raptor/raptor"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Route("GET", "/", "HomeController", "Root"),

		raptor.Scope("/movies",
			raptor.Route("GET", "/", "MoviesController", "Index"),
			raptor.Route("GET", "/:id", "MoviesController", "Show"),
		),
	)
}
