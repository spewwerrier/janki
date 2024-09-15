package db

import (
	"database/sql"
	"errors"
	"fmt"
	"janki/utils"
	"log"
)

// 1 user logins using username an password
// 2 user authenticates using cookie
// 3 get user information

func (db *Database) RetriveUserIdFromCredentials(username string, password string) (int, error) {
	query := "select id from users where username = $1"
	hashed_password, err := db.RetriveHashedPassword(username)
	if err == sql.ErrNoRows {
		return -1, err
	}
	if utils.CheckHash(hashed_password, password) {
		fmt.Println("password matches")
		result, err := db.db.Query(query, username)
		if err != nil {
			log.Panic(err)
			return -1, err
		}
		var id int
		var i int
		for result.Next() {
			err = result.Scan(&id)
			if err != nil {
				return -1, err
			}
			i++
		}
		if i == 0 {
			fmt.Println("no user exists with same name")
			return -1, sql.ErrNoRows
		}
		return id, nil
	}
	return -1, errors.New("cannot get user id")
}

func (db *Database) RetriveUserSession(username string, password string) (string, error) {
	query := "select id from users where username = $1 and password = $2"
	result, err := db.db.Query(query, username, utils.Hash(password))
	if err != nil {
		log.Panic(err)
		return "", err
	}
	var i int
	for result.Next() {
		i++
	}
	return "", err
}

func (db *Database) RetriveHashedPassword(username string) (string, error) {
	query := "select password from users where username = $1"
	result, err := db.db.Query(query, username)
	if err != nil {
		log.Panic(err)
		return "", err
	}
	var password string
	var i int
	for result.Next() {
		err = result.Scan(&password)
		if err != nil {
			return "", err
		}
		i++
	}
	if i == 0 {
		fmt.Println("user may not exists")
		return "", sql.ErrNoRows
	}
	return password, nil
}

func (db *Database) RetriveUserIdFromSession(cookie string) (int, error) {
	return -1, nil
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
