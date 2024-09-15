package db

// 1 user logins using username an password
// 2 user authenticates using cookie
// 3 user requests knobs using cookie
// 4 user requests knob descriptions using cookie and id
// if cookie->ref users(id) != knob->ref users(id) then we don't send the knob

func (db *Database) GetUserIdFromCredentials(username string, password string) error {
	return nil
}

func (db *Database) GetUserCookie(username string, password string) error {
	return nil
}

func (db *Database) GetUserIdFromCookie(cookie string) error {
	return nil
}

func (db *Database) GetUserKnobs(cookie string) error {
	return nil
}

func (db *Database) GetKnobId(cookie string) error {
	return nil
}

func (db *Database) GetKnobDescriptions(cookie string) error {
	return nil
}
