package db

import (
	"fmt"
	"log"

	"janki/jlog"

	"github.com/jackc/pgx/v5"
)

// takes api key and returns a knob
// uses RetriveUserIdFromSession which uses 1 sql query
// and does a select query
// this overall uses 2 sql queries
func (db *Database) GetUserKnobs(api_key string) ([]Knob, error) {
	id, err := db.RetriveUserIdFromApi(api_key)
	if err != nil {
		db.log.Error("GetUserKnob failed to get user id")
		return nil, err
	}
	query := "select knob_name,creation,ispublic, identifier from knobs inner join knobdescriptions on knobdescriptions.knob_id = knobs.id where knobs.user_id = $1 order by knobs.creation desc"
	result, err := db.Query(query, id)
	if err != nil {
		db.log.Error(err.Error())
		return nil, err
	}
	var i int
	var knob Knob
	var knobs []Knob
	fmt.Println(query, id)
	for result.Next() {
		i++
		result.Scan(&knob.KnobName, &knob.Creation, &knob.IsPublic, &knob.Identifier)
		knobs = append(knobs, knob)
	}
	if i < 1 {
		db.log.Warning("GetUserKnobs no knob exists")
		return nil, jlog.ErrNoKnobExists
	}
	return knobs, nil
}

// returns ErrNoKnobExists indicating new knob can be made == 0
// returns ErrKnobAlreadyExists indicating new knob cannot be made (we should never hit this) > 1
// return knobId if there is already a knob == 1
func (db *Database) GetKnobId(api_key string, knob_name string) (int, error) {
	id, err := db.GetUserId(api_key)
	if err != nil {
		return -1, err
	}
	query := "select id from knobs where user_id = $1 and knob_name = $2"
	result, err := db.Query(query, id, knob_name)
	if err != nil {
		return -1, err
	}
	var i int
	var knobId int
	for result.Next() {
		i++
		result.Scan(&knobId)
	}
	if i == 0 {
		return -1, jlog.ErrNoKnobExists
	}
	if i > 1 {
		return -1, jlog.ErrKnobAlreadyExists
	}
	return knobId, nil
}

func (db *Database) GetKnobIdFromIdentifier(identifier string) (int, error) {
	query := "select id from knobs where identifier = $1"
	result, err := db.Query(query, identifier)
	if err != nil {
		return -1, err
	}
	var i int
	var knobId int
	for result.Next() {
		i++
		result.Scan(&knobId)
	}
	if i == 0 {
		return -1, jlog.ErrNoKnobExists
	}
	if i > 1 {
		return -1, jlog.ErrKnobAlreadyExists
	}
	return knobId, nil
}

// TODO: if api key is given then it returns all public and private knobs else get only public knob
// we want error to be ignored
func (db *Database) GetKnobDescriptions(api string, identifier string) (KnobDescription, error) {
	knob := KnobDescription{}
	id, err := db.GetKnobIdFromIdentifier(identifier)
	if err != nil {
		db.log.Error("GetKnobDescriptions failed to get knob id from identifier")
		return knob, err
	}

	user_id, err := db.GetUserId(api)
	var result pgx.Rows

	query := "select * from knobs where user_id = $1 and identifier = $2"
	rows := db.QueryRow(query, user_id, identifier)
	err = rows.Scan()
	// if there is no error it means the original user is asking for the knob so we give the knob even if its private
	if err != nil {
		query = "select knobs.knob_name, knobs.identifier, description, topics, todo, tor, refs, urls, ques, suggestions, knobs.ispublic from knobdescriptions inner join knobs on knobs.id = knobdescriptions.id where knob_id = $1 and ispublic = true"
		result, err = db.Query(query, id)
		if err != nil {
			db.log.Error("GetKnobDescriptions failed to execute query")
			return knob, err
		}
	} else {
		query = "select knobs.knob_name, knobs.identifier, description, topics, todo, tor, refs, urls, ques, suggestions, knobs.ispublic from knobdescriptions inner join knobs on knobs.id = knobdescriptions.id  inner join users on  knobs.user_id = users.id where knob_id = $1 and user_id = $2"
		result, err = db.Query(query, id)
		if err != nil {
			db.log.Error("GetKnobDescriptions failed to execute query")
			return knob, err
		}
	}

	for result.Next() {
		err = result.Scan(&knob.Knob.KnobName, &knob.Knob.Identifier, &knob.Description, &knob.Topics, &knob.Todo, &knob.Tor, &knob.Refs, &knob.Urls, &knob.Ques, &knob.Suggestions, &knob.Knob.IsPublic)
		if err != nil {
			log.Panic(err)
		}
	}
	return knob, nil
}
