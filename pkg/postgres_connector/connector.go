package postgres_connector

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Connector struct {
	*sql.DB
}

func NewConnector(cfg Postgres) (*Connector, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
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
		DB: db,
	}, nil
}

func (d *Connector) Close() error {
	return d.Close()
}
