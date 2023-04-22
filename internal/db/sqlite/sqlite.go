package sqlite

import (
	"ProjectIdeas/monolith/internal/db"
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName = "index.db"
	dbPath = "/Users/jayantshelke/go/src/ProjectIdeas/monolith/internal/db/sqlite/"
)

func New(ctx context.Context) (db.DBer, error) {
	d, err := sql.Open("sqlite3", dbPath+dbName)
	if err != nil {
		return nil, err
	}

	if err := d.Ping(); err != nil {
		return nil, err
	}

	return &DB{d}, nil
}

type DB struct {
	db *sql.DB
}
