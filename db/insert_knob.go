package db

import (
	jankilog "janki/logs"
)

func (db *Database) CreateNewKnob(session_key string, knob Knob) error {
	id, err := db.RetriveUserIdFromSession(session_key)
	if err != nil {
		return err
	}

	_, err = db.GetKnobId(session_key, knob.KnobName)
	if err != jankilog.ErrNoKnobExists {
		return jankilog.ErrKnobExists
	}

	// // TODO: check this
	// if err := db.GetUserKnobs(session_key); err != jankilog.ErrNoKnobExists {
	// 	return errors.New("failed to create new knob. Knob already exists")
	// }
	// _, err = db.GetKnobId(session_key, knob.KnobName)
	// if err != jankilog.ErrNoKnobExists {
	// 	return errors.New("knob already exists")
	// }

	query := "insert into knobs (user_id, knob_name, ispublic) values ($1, $2, $3)"
	_, err = db.db.Exec(query, id, knob.KnobName, knob.IsPublic)
	if err != nil {
		return err
	}

	knobId, err := db.GetKnobId(session_key, knob.KnobName)
	if err != nil {
		return err
	}

	query = "insert into knobdescriptions (knob_id, description) values ($1, $2)"
	_, err = db.db.Exec(query, knobId, "this is a sample knob")
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) CreateKnobDescriptions(session_key string) error {
	return nil
}

func (db *Database) UpdateKnob(session_key string) error {
	return nil
}

func (db *Database) UpdateKnobDescriptions(cookie string) error {
	return nil
}

func (db *Database) DeleteKnob(cookie string) error {
	return nil
}
