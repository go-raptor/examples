package migrate

import (
	"github.com/go-raptor/examples/api/app/models"
	"github.com/go-raptor/raptor"
)

func AddMovies(db *raptor.DB) error {
	return db.AutoMigrate(&models.Movie{})
}
