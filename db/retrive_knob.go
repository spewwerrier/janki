package db

import (
	"fmt"

	"janki/jlog"
)

// 3 user requests knobs using cookie
// 4 user requests knob descriptions using cookie and id
// if cookie->ref users(id) != knob->ref users(id) then we don't send the knob

func (db *Database) GetUserKnobs(session_key string) ([]Knob, error) {
	id, err := db.RetriveUserIdFromSession(session_key)
	if err != nil {
		return nil, err
	}
	query := "select knob_name,creation,ispublic from knobs inner join knobdescriptions on knobdescriptions.knob_id = knobs.id where knobs.user_id = $1"
	result, err := db.raw.Query(query, id)
	var i int
	var knob Knob
	var knobs []Knob
	fmt.Println(query, id)
	for result.Next() {
		i++
		result.Scan(&knob.KnobName, &knob.Creation, &knob.IsPublic)
		knobs = append(knobs, knob)
	}
	if i < 1 {
		return nil, jlog.ErrNoKnobExists
	}
	return knobs, nil
}

func (db *Database) GetKnobId(session_key string, knob_name string) (int, error) {
	id, err := db.GetUserId(session_key)
	if err != nil {
		return -1, err
	}
	query := "select id from knobs where user_id = $1 and knob_name = $2"
	result, err := db.raw.Query(query, id, knob_name)
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

func (db *Database) GetKnobDescriptions(cookie string) error {
	return nil
}
