package db

import (
	"errors"
	"fmt"
)

// TODO: check if api and identifier belongs to same user
func (db *Database) UpdateKnobDescriptions(api_key string, identifier string, key string, value string) error {
	isUserCorrect, err := db.AuthorizeUserKnob(api_key, identifier)
	if err != nil {
		db.log.Error("UpdateKnobDescriptions failed to verify user")
		return err
	}
	if !isUserCorrect {
		return errors.New("User is not authorized to edit the knob")
	}

	knob_id, err := db.GetKnobIdFromIdentifier(identifier)
	if err != nil {
		db.log.Error("UpdateKnobDescriptions failed to get knob id")
		return err
	}

	var query string
	query = fmt.Sprintf("UPDATE knobdescriptions  SET %s = $1  WHERE knob_id = $2", key)
	fmt.Println(query, value, knob_id)
	_, err = db.Execute(query, value, knob_id)
	if err != nil {
		db.log.Error("UpdateUser: Failed to update user " + err.Error())
		return err
	}

	return nil
}

func (db *Database) UpdateKnobPublic(api_key string, identifier string, ispublic string) error {
	knob_id, err := db.GetKnobIdFromIdentifier(identifier)
	if err != nil {
		db.log.Error("UpdateKnob failed to get knob id")
		return err
	}
	query := ("UPDATE knobs SET ispublic = $1  WHERE id = $2")
	fmt.Println(query, ispublic, knob_id)
	_, err = db.Execute(query, ispublic, knob_id)
	if err != nil {
		db.log.Error("UpdateKnobPublic: failed to update knob public property " + err.Error())
		return err
	}
	return nil
}
