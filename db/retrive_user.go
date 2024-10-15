package db

import (
	"errors"

	"janki/jlog"
	"janki/utils"
)

// 1 user logins using username an password
// 2 user authenticates using cookie
// 3 get user information

func (db *Database) RetriveUserIdFromCredentials(username string, password string) (int, error) {
	query := "select id from users where username = $1"
	hashed_password, err := db.RetriveHashedPassword(username)
	if err == jlog.ErrApiUserNoExist {
		db.log.Warning("RetriveUserIdFromCredentaials " + err.Error())
		return -1, err
	}
	if utils.CheckHash(hashed_password, password) {
		result, err := db.raw.Query(query, username)
		if err != nil {
			db.log.Error("RetriveUserIdFromCredentials " + err.Error())
			return -1, jlog.ErrDbQueryError
		}
		var id int
		var i int
		for result.Next() {
			_ = result.Scan(&id)
			i++
		}
		if i == 0 {
			db.log.Warning("RetriveUserIdFromCredentaials: No user exists with name " + username)
			return -1, jlog.ErrApiUserNoExist
		}
		return id, nil
	}
	return -1, errors.New("cannot get user id")
}

func (db *Database) RetriveUserApi(username string, password string) (string, error) {
	id, err := db.RetriveUserIdFromCredentials(username, password)
	if err != nil {
		return "", err
	}
	query := "select session_key from sessions where user_id = $1"
	result, err := db.raw.Query(query, id)
	if err != nil {
		db.log.Error("RetriveUserApi " + err.Error())
		return "", err
	}
	var i int
	var session_key string
	for result.Next() {
		_ = result.Scan(&session_key)
		i++
	}
	return session_key, err
}

func (db *Database) RetriveHashedPassword(username string) (string, error) {
	query := "select password from users where username = $1"
	result, err := db.raw.Query(query, username)
	if err != nil {
		db.log.Error("RetriveHashedPassword " + err.Error())
		return "", jlog.ErrDbQueryError
	}
	var password string
	var i int
	for result.Next() {
		_ = result.Scan(&password)
		i++
	}
	if i == 0 {
		db.log.Warning("RetriveHashedPassword user does not exists")
		return "", jlog.ErrApiUserNoExist
	}
	return password, nil
}

func (db *Database) RetriveUserIdFromApi(api_key string) (int, error) {
	query := "select user_id from sessions where session_key = $1"
	result, err := db.raw.Query(query, api_key)
	if err != nil {
		db.log.Error("RetriveUserIdFromApi failed to query the database")
		return -1, err
	}
	var id int
	for result.Next() {
		_ = result.Scan(&id)
	}
	return id, nil
}

func (db *Database) CheckDuplicateknobs(session_key string, knob_name string) error {
	// get a knob with knob_name and id = session_key
	// if exists then duplicate

	return nil
}

func (db *Database) RetriveUser(cookie string) (UsersDetails, error) {
	u := UsersDetails{
		Info: Info{
			Name:     "Aagaman",
			Password: "lkdfsjasd8asdkljf;akjl",
		},
		Description: Descriptions{
			Creation:       120948,
			Image_url:      "https://example.com",
			Description:    "This is a dummy account",
			Existing_knobs: 129,
		},
		Session: Session{
			Cookie_string: "9080reujidskasdh780oguij",
			Creation:      3600,
			User_id:       348907,
		},
	}
	return u, nil
}
