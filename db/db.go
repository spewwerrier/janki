package db

import (
	"context"
	"database/sql"
	"fmt"

	"janki/jlog"
)

type Database struct {
	raw *sql.DB
	ctx context.Context
	log jlog.Jlog
}

func (db *Database) Execute(query string, args ...interface{}) (sql.Result, error) {
	db.log.Info(fmt.Sprintf("executing query: %s with args: %v", query, args))
	result, err := db.raw.Exec(query, args...)
	if err != nil {
		db.log.Error(fmt.Sprintf("Execute failed: %s, error: %v", query, err))
		return nil, err
	}
	return result, nil
}

func (db *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	db.log.Info(fmt.Sprintf("querying query: %s with args: %v", query, args))
	rows, err := db.raw.Query(query, args...)
	if err != nil {
		db.log.Error(fmt.Sprintf("Query failed: %s, error: %v", query, err))
		return nil, err
	}
	return rows, nil
}

func (db *Database) QueryRow(query string, args ...interface{}) (*sql.Row, error) {
	db.log.Info(fmt.Sprintf("querying query: %s with args: %v", query, args))
	row := db.raw.QueryRow(query, args...)
	if row.Err() != nil {
		db.log.Error(fmt.Sprintf("Query failed: %s, error: %v", query, row.Err()))
		return nil, row.Err()
	}
	return row, nil
}
