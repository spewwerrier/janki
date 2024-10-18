package db

import "fmt"

// TODO: check if api and identifier belongs to same user
func (db *Database) UpdateKnob(api_key string, identifier string, key string, value string) error {
	knob_id, err := db.GetKnobIdFromIdentifier(identifier)
	if err != nil {
		db.log.Error("UpdateKnob failed to get knob id")
		return err
	}
	var query string
	if key == "description" || key == "image_url" {
		query = fmt.Sprintf("UPDATE knobdescriptions  SET %s = $1  WHERE knob_id = $2", key)
	} else {
		query = fmt.Sprintf("UPDATE knobdescriptions  SET %s = %s || ARRAY[$1]  WHERE knob_id = $2", key, key)
	}
	_, err = db.raw.Exec(query, value, knob_id)
	if err != nil {
		db.log.Error("UpdateUser: Failed to update user " + err.Error())
		return err
	}
	return nil
}
