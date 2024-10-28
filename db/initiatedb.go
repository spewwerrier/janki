package db

import (
	"context"
	"fmt"
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
	_, err = d.Execute("create table if not exists Knobs (id serial primary key, user_id integer references Users(id) on delete cascade, knob_name text, description text, creation timestamp default current_timestamp not null, forkof text, ispublic bool, identifier text)")
	if err != nil {
		return err
	}

	_, err = d.Execute("create table if not exists KnobReferences (id serial primary key, knob_id integer references Knobs(id) on delete cascade, Refs text)")
	if err != nil {
		return err
	}
	_, err = d.Execute("create table if not exists KnobTopics (id serial primary key, knob_id integer references Knobs(id) on delete cascade, Topics text)")
	if err != nil {
		return err
	}

	_, err = d.Execute("create table if not exists KnobThingsToRead (id serial primary key, knob_id integer references Knobs(id) on delete cascade, ThingsToRead text)")
	if err != nil {
		return err
	}
	_, err = d.Execute("create table if not exists KnobUrls (id serial primary key, knob_id integer references Knobs(id) on delete cascade, Urls text)")
	if err != nil {
		return err
	}
	_, err = d.Execute("create table if not exists KnobQuestions (id serial primary key, knob_id integer references Knobs(id) on delete cascade, Questions text)")
	if err != nil {
		return err
	}
	_, err = d.Execute("create table if not exists KnobSuggestions (id serial primary key, knob_id integer references Knobs(id) on delete cascade, Suggestions text)")
	if err != nil {
		return err
	}
	_, err = d.Execute("create table if not exists KnobTodo (id serial primary key, knob_id integer references Knobs(id) on delete cascade, Todo text)")
	if err != nil {
		return err
	}

	// api
	_, err = d.Execute("create table if not exists Api (id serial primary key, api_key text not null, creation timestamp default current_timestamp not null, user_id integer references Users(id) on delete cascade)")
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) CleanDb() error {
	tables := []string{"api", "knobdescriptions", "knobquestions", "knobreferences", "knobsuggestions", "knobthingstoread", "knobtodo", "knobtopics", "knoburls", "sessions", "usersdescriptions", "knobs", "users"}
	for _, i := range tables {
		query := fmt.Sprintf("drop table if exists %s", i)
		_, err := d.Execute(query)
		if err != nil {
			return err
		}
	}
	return nil
}
