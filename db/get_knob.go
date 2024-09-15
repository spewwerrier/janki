package db

// 3 user requests knobs using cookie
// 4 user requests knob descriptions using cookie and id
// if cookie->ref users(id) != knob->ref users(id) then we don't send the knob

func (db *Database) GetUserKnobs(cookie string) error {
	return nil
}

func (db *Database) GetKnobId(cookie string) error {
	return nil
}

func (db *Database) GetKnobDescriptions(cookie string) error {
	return nil
}
