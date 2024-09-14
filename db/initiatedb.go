package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func NewConnection(connection string) *Database {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Panic(err)
	}
	return &Database{
		db: db,
	}
}

func (d *Database) Create_db() error {
	_, err := d.db.Exec("create table if not exists Users (id serial, name text, password text)")
	if err != nil {
		log.Panic(err)
	}
	return nil
}
