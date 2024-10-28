package db

import (
	"errors"
	"fmt"

	"janki/jlog"
	"janki/utils"
)

// key is the table you want to update and value is the update parameter
func (db *Database) InsertKnob(key string, value string, identifier string, api_key string) error {
	authorized, err := db.AuthorizeUserKnob(api_key, identifier)
	if err != nil {
		db.log.Error("InsertKnob failed to authorize the error")
		return err
	}
	if !authorized {
		return errors.New("User is not authorized")
	}
	// if key does not exists in the table we throw error
	err = jlog.ErrKnobItemDoesNotExists
	for _, v := range KnobItemTable {
		if key == v {
			err = nil
		}
	}
	if err != nil {
		return err
	}

	knobId, err := db.GetKnobIdFromIdentifier(identifier)
	if err != nil {
		db.log.Error("InsertKnob failed to retrive user id from api")
		return err
	}

	rawQuery := fmt.Sprintf("insert into %s values (DEFAULT, $1, $2)", key)
	_, err = db.Execute(rawQuery, knobId, value)
	if err != nil {
		return nil
	}
	return nil
}

// creates new knob using the api key
func (db *Database) CreateNewKnob(api_key string, knob Knob) (string, error) {
	id, err := db.RetriveUserIdFromApi(api_key)
	fmt.Println(id)
	if err != nil {
		db.log.Error("CreateNewKnob failed to retrive user id from api")
		return "", err
	}

	_, err = db.GetKnobId(api_key, knob.KnobName)
	// we dont want any knob of same name to already exist
	// so if we dont find error of KnobNotFound then there is alerady
	// an existing knob with same name :(((
	if err != jlog.ErrNoKnobExists {
		return "", jlog.ErrKnobAlreadyExists
	}

	knob.Identifier = utils.GenerateIdentifier(int64(id))

	query := "insert into knobs (user_id, knob_name, ispublic, identifier, forkof, description) values ($1, $2, $3, $4, $5, $6)"
	_, err = db.Execute(query, id, knob.KnobName, knob.IsPublic, knob.Identifier, knob.ForkOf, knob.Description)
	if err != nil {
		db.log.Error("CreateNewKnob failed to create knob: " + query)
		return "", err
	}

	// knobId, err := db.GetKnobId(api_key, knob.KnobName)
	// if err != nil {
	// 	return "", err
	// }

	return knob.Identifier, nil
}

// TODO: match api and knob and if they same delete knob
func (db *Database) DeleteKnob(api string, knob_id string) error {
	return nil
}

func (db *Database) ForkKnob(api string, identifier string) error {
	knob, err := db.GetKnobDescriptions(api, identifier)
	if err != nil {
		db.log.Error("ForkKnob failed to get user knobs")
		return err
	}
	fmt.Println(api, identifier, knob)

	knob.ForkOf = knob.Identifier
	_, err = db.CreateNewKnob(api, knob)
	if err != nil {
		db.log.Error("ForkKnob failed to create new knob " + err.Error())
		return err
	}
	return nil
}
