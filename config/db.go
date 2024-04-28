package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

type Options struct {
	DatabaseURL string
}

func New(cfg Options) (*Postgres, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	
	err = db.Ping()
	
	if err != nil {
		log.Fatal(err)
	}
	
	return &Postgres{Db: db}, nil
}
