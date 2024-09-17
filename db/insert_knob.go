package db

// 5 user creates new knob
// 6 user updates their knob
// 7 user updates their knob descriptions
// 8 user deletes their knob

func (db *Database) CreateNewKnob(session_key string, knob_name string) error {
	id, err := db.RetriveUserIdFromSession(session_key)
	if err != nil {
		return err
	}
	query := "insert into knobs (user_id, knob_name) values ($1, $2)"
	_, err = db.db.Exec(query, id, knob_name)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateKnob(session_key string) error {
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
