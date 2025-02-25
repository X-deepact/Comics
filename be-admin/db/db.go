package db

import (
	"context"
	"gorm.io/gorm"
)

// The DBTX interface makes it easy to switch between different database sessions or
// transactions without changing the logic of your application.
// For example, you can pass a transactional *gorm.DB or
// a regular database connection (*gorm.DB) depending on the context of the operation.
// If you intend to support both *gorm.DB and *sql.Tx,

type DBTX interface {
	WithContext(ctx context.Context) *gorm.DB
	//ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	//QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

type Queries struct {
	db DBTX
}

// New creates a new Queries instance
func New(db DBTX) *Queries {
	return &Queries{db: db}
}
