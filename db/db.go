package db

import "database/sql"

type Database struct {
	db *sql.DB
}
