package migrate

import (
	"github.com/go-raptor/examples/api/app/models"
	"github.com/go-raptor/raptor"
)

func SeedMovies(db *raptor.DB) error {
	movies := models.Movies{
		models.Movie{
			Title:       "The Shawshank Redemption",
			Description: "A gripping tale of hope and redemption as a man finds solace and purpose in the confines of a prison.",
			Year:        1994,
		},
		models.Movie{
			Title:       "The Godfather",
			Description: "A powerful crime family saga that follows the transformation of Michael Corleone from a reluctant family outsider to a ruthless mafia boss.",
			Year:        1972,
		},
	}
	return db.CreateInBatches(&movies, 2).Error
}
