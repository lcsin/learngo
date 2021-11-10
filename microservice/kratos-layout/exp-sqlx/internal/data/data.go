package data

import (
	"exp-sqlx/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewArticleRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *sqlx.DB
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	db, err := sqlx.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	cleanup := func() {
		log.Info("message", "closing the data resources")
		if err := db.Close(); err != nil {
			log.Error("failed close connection to mysql: %v", err)
		}
	}

	return &Data{
		db:  db,
		log: log,
	}, cleanup, nil
}
