package db

import (
	"errors"
	jankilog "janki/logs"
	"janki/utils"
)

// 1 user logins using username an password
// 2 user authenticates using cookie
// 3 get user information

func (db *Database) RetriveUserIdFromCredentials(username string, password string) (int, error) {
	query := "select id from users where username = $1"
	hashed_password, err := db.RetriveHashedPassword(username)
	if err == jankilog.ErrApiUserNoExist {
		db.log.Warning("db: " + err.Error())
		return -1, err
	}
	if utils.CheckHash(hashed_password, password) {
		result, err := db.db.Query(query, username)
		if err != nil {
			db.log.Error(err.Error())
			return -1, jankilog.ErrDbQueryError
		}
		var id int
		var i int
		for result.Next() {
			_ = result.Scan(&id)
			i++
		}
		if i == 0 {
			db.log.Warning("db: No user exists with name " + username)
			return -1, jankilog.ErrApiUserNoExist
		}
		return id, nil
	}
	return -1, errors.New("db: cannot get user id")
}

func (db *Database) RetriveUserSession(username string, password string) (string, error) {
	id, err := db.RetriveUserIdFromCredentials(username, password)
	if err != nil {
		return "", err
	}
	query := "select session_key from sessions where user_id = $1"
	result, err := db.db.Query(query, id)
	if err != nil {
		db.log.Error("db: " + err.Error())
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
	result, err := db.db.Query(query, username)
	if err != nil {
		db.log.Error("db: " + err.Error())
		return "", jankilog.ErrDbQueryError
	}
	var password string
	var i int
	for result.Next() {
		_ = result.Scan(&password)
		i++
	}
	if i == 0 {
		return "", jankilog.ErrApiUserNoExist
	}
	return password, nil
}

func (db *Database) RetriveUserIdFromSession(session_key string) (int, error) {
	query := "select user_id from users where session_key = $1"
	result, err := db.db.Query(query, session_key)
	if err != nil {
		return -1, nil
	}
	var id int
	for result.Next() {
		_ = result.Scan(&id)
	}
	return id, nil
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
