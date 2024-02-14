package db

import (
	"github.com/go-raptor/examples/api/db/migrate"
	"github.com/go-raptor/raptor"
)

func Migrations() raptor.Migrations {
	return raptor.Migrations{
		1: migrate.AddMovies,
		2: migrate.SeedMovies,
	}
}
