package db

import (
	"fmt"

	"janki/jlog"
	"janki/utils"
)

// creates new knob using the api key
func (db *Database) CreateNewKnob(api_key string, knob KnobDescription) (string, error) {
	id, err := db.RetriveUserIdFromApi(api_key)
	fmt.Println(id)
	if err != nil {
		db.log.Error("CreateNewKnob failed to retrive user id from api")
		return "", err
	}

	_, err = db.GetKnobId(api_key, knob.Knob.KnobName)
	// we dont want any knob of same name to already exist
	// so if we dont find error of KnobNotFound then there is alerady
	// an existing knob with same name :(((
	if err != jlog.ErrNoKnobExists {
		return "", jlog.ErrKnobAlreadyExists
	}

	knob.Knob.Identifier = utils.GenerateIdentifier(int64(id))

	query := "insert into knobs (user_id, knob_name, ispublic, identifier, forkof) values ($1, $2, $3, $4, $5)"
	_, err = db.Execute(query, id, knob.Knob.KnobName, knob.Knob.IsPublic, knob.Knob.Identifier, knob.Knob.ForkOf)
	if err != nil {
		db.log.Error("CreateNewKnob failed to create knob: " + query)
		return "", err
	}

	knobId, err := db.GetKnobId(api_key, knob.Knob.KnobName)
	if err != nil {
		return "", err
	}

	query = "insert into knobdescriptions (knob_id, topics, todo, tor, refs, urls, ques, description, suggestions) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err = db.Execute(query, knobId, knob.Topics, knob.Todo, knob.Tor, knob.Refs, knob.Urls, knob.Ques, knob.Description, knob.Suggestions)
	if err != nil {
		db.log.Error("CreateNewKnob failed to create knob descriptions: " + query)
		return "", err
	}

	return knob.Knob.Identifier, nil
}

// TODO
func (db *Database) DeleteKnob(api string) error {
	return nil
}

func (db *Database) ForkKnob(api string, identifier string) error {
	knob, err := db.GetKnobDescriptions(api, identifier)
	if err != nil {
		db.log.Error("ForkKnob failed to get user knobs")
		return err
	}
	fmt.Println(api, identifier, knob)

	knob.Knob.ForkOf = knob.Knob.Identifier
	_, err = db.CreateNewKnob(api, knob)
	if err != nil {
		db.log.Error("ForkKnob failed to create new knob " + err.Error())
		return err
	}
	return nil
}
