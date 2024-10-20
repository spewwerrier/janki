package db

import (
	"context"
	"fmt"
	"time"

	"janki/jlog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	raw *pgxpool.Pool
	ctx context.Context
	log jlog.Jlog
}

func (db *Database) Execute(query string, args ...interface{}) (pgconn.CommandTag, error) {
	ctx, cancel := context.WithTimeout(db.ctx, time.Second*10)
	defer cancel()

	db.log.Info(fmt.Sprintf("executing query: %s with args: %v", query, args))
	result, err := db.raw.Exec(ctx, query, args...)
	if err != nil {
		db.log.Error(fmt.Sprintf("Execute failed: %s, error: %v", query, err))
		return pgconn.CommandTag{}, err
	}
	return result, nil
}

func (db *Database) Query(query string, args ...interface{}) (pgx.Rows, error) {
	ctx, cancel := context.WithTimeout(db.ctx, time.Second*10)
	defer cancel()

	db.log.Info(fmt.Sprintf("querying query: %s with args: %v", query, args))
	rows, err := db.raw.Query(ctx, query, args...)
	if err != nil {
		db.log.Error(fmt.Sprintf("Query failed: %s, error: %v", query, err))
		return nil, err
	}
	return rows, nil
}

func (db *Database) QueryRow(query string, args ...interface{}) pgx.Row {
	ctx, cancel := context.WithTimeout(db.ctx, time.Second*10)
	defer cancel()
	db.log.Info(fmt.Sprintf("querying queryrow: %s with args: %v", query, args))
	row := db.raw.QueryRow(ctx, query, args...)

	return row
}
