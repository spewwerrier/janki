package db

import (
	"database/sql"
	"errors"
	"janki/utils"
	"log"
	"strconv"
	"time"
)

// 1 user registers their username and password
// 2 user updates their username and password
// 3 user adds their user_description field
// 4 user deletes their account
// 5 generate cookie
// 6 check if duplicate user exists

func (db *Database) CreateNewUser(username string, password string) (string, error) {
	does, err := db.CheckDuplicateUser(username)
	if err != nil {
		panic(err)
	}
	if !does {
		query := "insert into users (username, password) values ($1, $2)"
		_, err := db.db.Exec(query, username, utils.Hash(password))
		if err != nil {
			log.Panic(err)
		}

		session_key, err := db.GenerateSessionKey(username, password)
		if err != nil {
			log.Panic(err)
			return "", err
		}
		return session_key, nil
	}
	return "", errors.New("duplicate user exists")
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

func (db *Database) DeleteAccount(cookie string) error {
	return nil
}

func (db *Database) GenerateSessionKey(username string, password string) (string, error) {
	session_key := utils.Hash(username + password + strconv.Itoa(time.Now().Nanosecond()))
	id, err := db.RetriveUserIdFromCredentials(username, password)
	if err != nil {
		log.Panic(err)
	}

	sql := "insert into sessions (session_key, user_id) values ($1, $2)"
	_, err = db.db.Exec(sql, session_key, id)
	if err != nil {
		log.Panic(err)
	}

	return session_key, nil
}

func (db *Database) CheckDuplicateUser(username string) (bool, error) {
	_, err := db.RetriveUserIdFromCredentials(username, "")
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}
