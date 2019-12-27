package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Datastore have all methods for database
type Datastore interface {
}

// DB init struct
type DB struct {
	*sqlx.DB

	ctx context.Context
}

// CreateDB connection to database
func CreateDB(host string, port int, user, password, database, sslmode string) (*DB, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)

	db, err := sqlx.ConnectContext(ctx, "postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		user,
		password,
		database,
		sslmode))
	if err != nil {
		return nil, fmt.Errorf("connect to database error: %+v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping to database error: %+v", err)
	}

	return &DB{db, ctx}, nil
}
