package db

import (
	"database/sql"
	jankilog "janki/logs"
	"log"

	_ "github.com/lib/pq"
)

func NewConnection(connection string, logfile string) *Database {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Panic(err)
	}
	logs := jankilog.NewLogger(logfile)
	return &Database{
		db:  db,
		log: logs,
	}
}

func (d *Database) Create_db() error {
	// Users
	_, err := d.db.Exec("create table if not exists Users (id serial primary key, username text not null, password text not null)")
	if err != nil {
		return err
	}

	// UsersDescription
	_, err = d.db.Exec("create table if not exists UsersDescriptions (user_id integer references Users(id) on delete cascade, creation timestamp default current_timestamp not null, image_url text, description text, existing_knobs int)")
	if err != nil {
		return err
	}

	// Knob
	_, err = d.db.Exec("create table if not exists Knobs (id serial primary key, user_id integer references Users(id) on delete cascade, creation timestamp default current_timestamp not null, forkof integer references Knobs(id), ispublic bool)")
	if err != nil {
		return err
	}

	// KnobDescriptions
	_, err = d.db.Exec("create table if not exists KnobDescriptions (knob_id integer references Knobs(id) on delete cascade, knob_name text, topics text[], todo text[], tor text[], refs text[], urls text[], ques text[], description text, suggestions text[])")
	if err != nil {
		return err
	}

	// Sessions
	_, err = d.db.Exec("create table if not exists Sessions (id serial primary key, session_key text not null, creation timestamp default current_timestamp not null, user_id integer references Users(id) on delete cascade)")
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) CleanDb() error {
	_, err := d.db.Exec("drop table if exists sessions")
	if err != nil {
		return err
	}
	_, err = d.db.Exec("drop table if exists knobdescriptions")
	if err != nil {
		return err
	}
	_, err = d.db.Exec("drop table if exists usersdescriptions")
	if err != nil {
		return err
	}
	_, err = d.db.Exec("drop table if exists knobs")
	if err != nil {
		return err
	}
	_, err = d.db.Exec("drop table if exists users")
	if err != nil {
		return err
	}
	return nil
}
