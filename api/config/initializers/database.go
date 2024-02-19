package initializers

import (
	"github.com/go-raptor/connector/sqlite"
	"github.com/go-raptor/examples/api/db"
	"github.com/go-raptor/raptor"
)

func Database() raptor.Database {
	return raptor.Database{
		Connector:  sqlite.NewSqliteConnector(),
		Migrations: db.Migrations(),
	}
}
