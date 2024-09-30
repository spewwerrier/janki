package db

import (
	jankilog "janki/logs"
)

// 3 user requests knobs using cookie
// 4 user requests knob descriptions using cookie and id
// if cookie->ref users(id) != knob->ref users(id) then we don't send the knob

func (db *Database) GetUserKnobs(session_key string) error {
	id, err := db.RetriveUserIdFromSession(session_key)
	if err != nil {
		return err
	}
	query := "select * from knobs where user_id = $1"
	result, err := db.db.Query(query, id)
	var i int
	for result.Next() {
		i++
	}
	if i < 1 {
		return jankilog.ErrNoKnobExists
	}
	return nil
}

func (db *Database) GetKnobId(session_key string, knob_name string) (int, error) {
	id, err := db.GetUserId(session_key)
	if err != nil {
		return -1, err
	}
	query := "select id from knobs where user_id = $1 and knob_name = $2"
	result, err := db.db.Query(query, id, knob_name)
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
		return -1, jankilog.ErrNoKnobExists
	}
	if i > 1 {
		return -1, jankilog.ErrKnobExists
	}
	return knobId, nil
}

func (db *Database) GetKnobDescriptions(cookie string) error {
	return nil
}
