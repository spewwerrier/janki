package db

// 1 user registers their username and password
// 2 user updates their username and password
// 3 user adds their user_description field
// 4 user creates new knob
// 5 user updates their knob
// 6 user updates their knob descriptions
// 7 user deletes their knob
// 8 user deletes their account

func (db *Database) CreateNewUser(username string, password string) error {
	return nil
}

func (db *Database) UpdateUser(cookie string) error {
	return nil
}

func (db *Database) CreateUserDescription(cookie string) error {
	return nil
}

func (db *Database) UpdateUserDescriptions(cookie string) error {
	return nil
}

func (db *Database) CreateNewKnob(cookie string) error {
	return nil
}

func (db *Database) UpdateKnob(cookie string) error {
	return nil
}

func (db *Database) CreateKnobDescriptions(cookie string) error {
	return nil
}

func (db *Database) UpdateKnobDescriptions(cookie string) error {
	return nil
}

func (db *Database) DeleteKnob(cookie string) error {
	return nil
}

func (db *Database) DeleteAccount(cookie string) error {
	return nil
}
