package db

import (
	"errors"
	jankilog "janki/logs"
	"janki/utils"
	"strconv"
	"time"
)

// 1 user registers their username and password
// 2 user updates their username and password
// 3 user adds their user_description field
// 4 user deletes their account
// 5 generate cookie
// 6 check if duplicate user exists

func (db *Database) CreateNewUser(username string, password string, image_url string, description string) (string, error) {
	does, err := db.CheckDuplicateUser(username)
	if err == jankilog.ErrApiMultipleUsers {
		return "", jankilog.ErrApiMultipleUsers
	}

	if does {
		return "", errors.New("duplicate user exists")
	}

	query := "insert into users (username, password) values ($1, $2)"
	_, err = db.db.Exec(query, username, utils.Hash(password))
	if err != nil {
		db.log.Warning(err.Error())
		return "", jankilog.ErrDbExecError
	}

	session_key, err := db.GenerateSessionKey(username, password)
	if err != nil {
		db.log.Warning(err.Error())
		return "", err
	}
	query = "insert into usersdescriptions (user_id, image_url, description) values ($1, $2, $3)"

	user_id, err := db.GetUserId(session_key)
	if err != nil {
		return "", err
	}
	_, err = db.db.Exec(query, user_id, image_url, description)
	if err != nil {
		return "", jankilog.ErrDbExecError
	}
	db.log.Info("db: inserted new user " + username)
	return session_key, nil
}

func (db *Database) UpdateUser(session_key string, image_url string, description string) error {

	query := "update usersdescriptions  set image_url = $1, description = $2 where user_id = $3"
	id, err := db.GetUserId(session_key)
	if err != nil {
		return err
	}
	_, err = db.db.Exec(query, image_url, description, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetUserId(session_key string) (int, error) {

	query := "select user_id from sessions where session_key = $1"
	result, err := db.db.Query(query, session_key)
	if err != nil {
		return -1, jankilog.ErrDbQueryError
	}

	var user_id int
	var i int
	for result.Next() {
		i++
		_ = result.Scan(&user_id)
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
	result, err := db.db.Query(query, user_id)
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
	_, err = db.db.Exec(query, user_id, image_url, description)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateUserDescriptions(cookie string) error {
	return nil
}

func (db *Database) DeleteAccount(cookie string) error {
	return nil
}

func (db *Database) GenerateSessionKey(username string, password string) (string, error) {
	session_key := utils.Hash(username + password + strconv.Itoa(time.Now().Nanosecond()))
	id, err := db.RetriveUserIdFromCredentials(username, password)
	if err != nil {
		return "", err
	}

	sql := "insert into sessions (session_key, user_id) values ($1, $2)"
	_, err = db.db.Exec(sql, session_key, id)
	if err != nil {
		db.log.Warning(err.Error())
		return "", jankilog.ErrDbExecError
	}

	return session_key, nil
}

func (db *Database) CheckDuplicateUser(username string) (bool, error) {
	_, err := db.RetriveUserIdFromCredentials(username, "")
	if err == jankilog.ErrApiUserNoExist {
		return false, nil
	}
	return true, jankilog.ErrApiMultipleUsers
}
