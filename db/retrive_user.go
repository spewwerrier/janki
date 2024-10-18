package db

import (
	"context"
	"errors"
	"fmt"
	"time"

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
		db.log.Error("RetriveUserApi failed to retrive user" + err.Error())
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
	ctx, cancel := context.WithTimeout(db.ctx, time.Second*5)
	defer cancel()
	query := "select user_id from sessions where session_key = $1"
	result, err := db.raw.QueryContext(ctx, query, api_key)
	// TODO: implement context so 5 seconds delay would be cancelled
	// time.Sleep(time.Second * 3)
	if err != nil {
		db.log.Error("RetriveUserIdFromApi failed to query the database" + err.Error())
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

func (db *Database) RetriveUser(api_key string) (UserDescription, error) {
	id, err := db.GetUserId(api_key)
	if err != nil {
		db.log.Error("RetriveUser falied to get user id")
		return UserDescription{}, nil
	}

	query := "select users.username, usersdescriptions.creation, image_url, description, usersdescriptions.creation, sessions.session_key, sessions.creation from users inner join usersdescriptions on users.id = usersdescriptions.user_id  inner join sessions on users.id = sessions.user_id where users.id = $1"
	rows := db.raw.QueryRow(query, id)
	if rows.Err() != nil {
		db.log.Error("RetriveUser failed to query the user information")
		return UserDescription{}, nil
	}
	u := UserDescription{}
	err = rows.Scan(
		&u.User.Name,
		&u.Creation,
		&u.Image_url,
		&u.Description,
		&u.Creation,
		&u.Session.ApiKey,
		&u.Session.Creation,
	)
	if err != nil {
		return u, err
	}
	fmt.Println(u)

	return u, nil
}
