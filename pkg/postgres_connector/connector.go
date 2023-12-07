package postgres_connector

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connector struct {
	db *sqlx.DB
}

func NewConnector(cfg Postgres) (*Connector, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"postgresql://root:%s@%s:%s/%s?sslmode=disable",
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Connector{
		db: db,
	}, nil
}

func (d *Connector) Close() error {
	return d.Close()
}

func (d *Connector) GetDB() *sqlx.DB {
	return d.db
}
