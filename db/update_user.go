package db

import "fmt"

func (db *Database) UpdateUser(api_key string, key string, value string) error {
	query := fmt.Sprintf("UPDATE usersdescriptions set %s = $1 where user_id  = $2", key)
	id, err := db.GetUserId(api_key)
	if err != nil {
		db.log.Error("UpdateUser: Failed to get user id " + err.Error())
		return err
	}
	_, err = db.Execute(query, value, id)
	fmt.Println(query, value, id)
	if err != nil {
		db.log.Error("UpdateUser: Failed to update user " + err.Error())
		return err
	}
	return nil
}
