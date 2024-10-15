package db

import (
	"errors"
	"fmt"
	"time"

	"janki/jlog"
	"janki/utils"
)

// creates new user and returns their api key
func (db *Database) CreateNewUser(username string, password string, image_url string, description string) (string, error) {
	does, err := db.CheckDuplicateUser(username)
	if err == jlog.ErrApiMultipleUsers {
		return "", jlog.ErrApiMultipleUsers
	}

	if does {
		return "", errors.New("duplicate user exists")
	}

	query := "insert into users (username, password) values ($1, $2)"
	_, err = db.raw.Exec(query, username, utils.HashBcrypt(password))
	if err != nil {
		db.log.Warning(err.Error())
		return "", jlog.ErrDbExecError
	}

	api_key, err := db.GenerateApiKey(username, password)
	if err != nil {
		db.log.Warning(err.Error())
		return "", err
	}
	query = "insert into usersdescriptions (user_id, image_url, description) values ($1, $2, $3)"

	user_id, err := db.GetUserId(api_key)
	if err != nil {
		return "", err
	}
	_, err = db.raw.Exec(query, user_id, image_url, description)
	if err != nil {
		return "", jlog.ErrDbExecError
	}
	db.log.Info("db: inserted new user " + username)
	return api_key, nil
}

// func (db *Database) UpdateUser(session_key string, image_url string, description string) error {
// 	query := "update usersdescriptions  set image_url = $1, description = $2 where user_id = $3"
// 	id, err := db.GetUserId(session_key)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = db.raw.Exec(query, image_url, description, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (db *Database) GetUserId(session_key string) (int, error) {
	query := "select user_id from sessions where session_key = $1"
	result, err := db.raw.Query(query, session_key)
	if err != nil {
		return -1, jlog.ErrDbQueryError
	}

	var user_id int
	var i int
	for result.Next() {
		i++
		_ = result.Scan(&user_id)
	}
	if i < 1 {
		return -1, jlog.ErrApiUserNoExist
	}
	if i > 1 {
		return user_id, errors.New("multiple user exists")
	}
	return user_id, nil
}

func (db *Database) CreateUserDescription(session_key string, image_url string, description string) error {
	user_id, err := db.GetUserId(session_key)
	if err != nil {
		return err
	}

	query := "select from usersdescriptions where user_id = $1"
	result, err := db.raw.Query(query, user_id)
	if err != nil {
		return err
	}
	var i int
	for result.Next() {
		i++
	}
	if i > 0 {
		return errors.New("descriptions already exists")
	}

	query = "insert into usersdescriptions (user_id, image_url, description) values ($1, $2, $3)"
	_, err = db.raw.Exec(query, user_id, image_url, description)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) RegenerateSessionKey(username string, password string) (string, error) {
	query := "delete from sessions where id = $1"
	id, err := db.RetriveUserIdFromCredentials(username, password)
	if err != nil {
		return "", err
	}
	_, err = db.raw.Exec(query, id)
	if err != nil {
		return "", err
	}
	session_key, err := db.GenerateApiKey(username, password)
	if err != nil {
		return "", err
	}
	return session_key, nil
}

func (db *Database) DeleteAccount(cookie string) error {
	return nil
}

func (db *Database) GenerateApiKey(username string, password string) (string, error) {
	api_key := utils.GenerateIdentifier(int64(time.Now().Nanosecond()))
	id, err := db.RetriveUserIdFromCredentials(username, password)
	if err != nil {
		return "", err
	}

	sql := "insert into sessions (session_key, user_id) values ($1, $2)"
	_, err = db.raw.Exec(sql, api_key, id)
	fmt.Println(api_key)
	if err != nil {
		db.log.Error("error while inserting session key in generatesessionkey\n" + err.Error())
		return "", jlog.ErrDbExecError
	}

	return api_key, nil
}

func (db *Database) CheckDuplicateUser(username string) (bool, error) {
	_, err := db.RetriveUserIdFromCredentials(username, "")
	if err == jlog.ErrApiUserNoExist {
		return false, nil
	}
	return true, jlog.ErrApiMultipleUsers
}
