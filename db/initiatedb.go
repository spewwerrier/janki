package db

import (
	"context"
	"log"

	"janki/jlog"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewConnection(connection string, logfile string) *Database {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, connection)
	if err != nil {
		log.Panic(err)
	}

	logs := jlog.NewLogger(logfile)
	return &Database{
		raw: conn,
		log: logs,
		ctx: ctx,
	}
}

func (d *Database) Close() {
	d.raw.Close()
}

func (d *Database) Create_db() error {
	// Users
	_, err := d.Execute("create table if not exists Users (id serial primary key, username text not null, password text not null)")
	if err != nil {
		return err
	}

	// UsersDescription
	_, err = d.Execute("create table if not exists UsersDescriptions (user_id integer references Users(id) on delete cascade, creation timestamp default current_timestamp not null, image_url text, description text, existing_knobs int)")
	if err != nil {
		return err
	}

	// Knob
	_, err = d.Execute("create table if not exists Knobs (id serial primary key, user_id integer references Users(id) on delete cascade, knob_name text, creation timestamp default current_timestamp not null, forkof text, ispublic bool, identifier text)")
	if err != nil {
		return err
	}

	// KnobDescriptions
	_, err = d.Execute("create table if not exists KnobDescriptions (id serial primary key, knob_id integer references Knobs(id) on delete cascade, topics text[], todo text[], tor text[], refs text[], urls text[], ques text[], description text, suggestions text[])")
	if err != nil {
		return err
	}

	// Sessions
	_, err = d.Execute("create table if not exists Sessions (id serial primary key, session_key text not null, creation timestamp default current_timestamp not null, user_id integer references Users(id) on delete cascade)")
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) CleanDb() error {
	_, err := d.Execute("drop table if exists sessions")
	if err != nil {
		return err
	}
	_, err = d.Execute("drop table if exists knobdescriptions")
	if err != nil {
		return err
	}
	_, err = d.Execute("drop table if exists usersdescriptions")
	if err != nil {
		return err
	}
	_, err = d.Execute("drop table if exists knobs")
	if err != nil {
		return err
	}
	_, err = d.Execute("drop table if exists users")
	if err != nil {
		return err
	}
	return nil
}
