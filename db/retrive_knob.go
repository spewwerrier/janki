package db

import (
	"fmt"
	"log"

	"janki/jlog"
)

// takes api key and returns a knob
// uses RetriveUserIdFromSession which uses 1 sql query
// and does a select query
// this overall uses 2 sql queries
func (db *Database) GetUserKnobs(api_key string) ([]Knob, error) {
	id, err := db.RetriveUserIdFromApi(api_key)
	if err != nil {
		return nil, err
	}
	query := "select knob_name,creation,ispublic, identifier from knobs inner join knobdescriptions on knobdescriptions.knob_id = knobs.id where knobs.user_id = $1 order by knobs.creation desc"
	result, err := db.Query(query, id)
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

func (db *Database) GetKnobDescriptions(api string, identifier string) (KnobDescription, error) {
	knob := KnobDescription{}
	id, err := db.GetKnobIdFromIdentifier(identifier)
	if err != nil {
		db.log.Error("GetKnobDescriptions failed to get knob id from identifier")
		return knob, err
	}
	query := "select description, topics, todo, tor, refs, urls, ques, suggestions from knobdescriptions where knob_id = $1"
	result, err := db.Query(query, id)
	if err != nil {
		db.log.Error("GetKnobDescriptions failed to execute query")
		return knob, err
	}
	for result.Next() {
		err = result.Scan(&knob.Description, &knob.Topics, &knob.Todo, &knob.Tor, &knob.Refs, &knob.Urls, &knob.Ques, &knob.Suggestions)
		if err != nil {
			log.Panic(err)
		}
	}
	return knob, nil
}
