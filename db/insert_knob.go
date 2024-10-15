package db

import (
	"fmt"

	"janki/jlog"
	"janki/utils"
)

// creates new knob using the api key
func (db *Database) CreateNewKnob(api_key string, knob Knob) error {
	id, err := db.RetriveUserIdFromApi(api_key)
	fmt.Println(id)
	if err != nil {
		return err
	}

	_, err = db.GetKnobId(api_key, knob.KnobName)
	// we dont want any knob of same name to already exist
	// so if we dont find error of KnobNotFound then there is alerady
	// an existing knob with same name :(((
	if err != jlog.ErrNoKnobExists {
		return jlog.ErrKnobAlreadyExists
	}

	knob.Identifier = utils.GenerateIdentifier(int64(id))

	query := "insert into knobs (user_id, knob_name, ispublic, identifier) values ($1, $2, $3, $4)"
	_, err = db.raw.Exec(query, id, knob.KnobName, knob.IsPublic, knob.Identifier)
	if err != nil {
		db.log.Error("CreateNewKnob failed to create knob: " + query)
		return err
	}

	knobId, err := db.GetKnobId(api_key, knob.KnobName)
	if err != nil {
		return err
	}

	query = "insert into knobdescriptions (knob_id, description) values ($1, $2)"
	_, err = db.raw.Exec(query, knobId, "this is a sample knob")
	if err != nil {
		db.log.Error("CreateNewKnob failed to create knob descriptions: " + query)
		return err
	}

	return nil
}

func (db *Database) CreateKnobDescriptions(session_key string) error {
	return nil
}

func (db *Database) UpdateKnobDescriptions(cookie string) error {
	return nil
}

func (db *Database) DeleteKnob(cookie string) error {
	return nil
}
