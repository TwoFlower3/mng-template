package models

import (
	"context"
	"fmt"
	"time"

	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// CreateDB connection to database
func TestCreateDB(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)

	sqlx.ConnectContext(ctx, "postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		"localhost",
		5432,
		"backend",
		"backend",
		"backend",
		"false"))
}
