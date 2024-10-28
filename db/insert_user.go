package db

import (
	"errors"
	"fmt"
	"time"

	"janki/jlog"
	"janki/utils"
)

// creates new user and returns their api key
func (db *Database) CreateNewUser(username string, password string) (string, error) {
	err := db.CheckDuplicateUser(username)
	if err != nil {
		db.log.Error("CreateNewUser duplicate user found")
		return "", jlog.ErrApiMultipleUsers
	}

	// if does {
	// 	return "", errors.New("duplicate user exists")
	// }

	query := "insert into users (username, password) values ($1, $2)"
	_, err = db.Execute(query, username, utils.HashBcrypt(password))
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
	_, err = db.Execute(query, user_id, "something", "something")
	if err != nil {
		return "", jlog.ErrDbExecError
	}
	db.log.Info("db: inserted new user " + username)
	return api_key, nil
}

// retrieves user id from the api key
func (db *Database) GetUserId(api_key string) (int, error) {
	query := "select user_id from api where api_key = $1"
	result, err := db.Query(query, api_key)
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

func (db *Database) RegenerateApiKey(username string, password string) (string, error) {
	query := "delete from api where id = $1"
	id, err := db.RetriveUserIdFromCredentials(username, password)
	if err != nil {
		return "", err
	}
	_, err = db.Execute(query, id)
	if err != nil {
		return "", err
	}
	api_key, err := db.GenerateApiKey(username, password)
	if err != nil {
		return "", err
	}
	return api_key, nil
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

	sql := "insert into api (api_key, user_id) values ($1, $2)"
	_, err = db.Execute(sql, api_key, id)
	fmt.Println(api_key)
	if err != nil {
		db.log.Error("GenerateApiKey error while inserting api key in api key\n" + err.Error())
		return "", jlog.ErrDbExecError
	}

	return api_key, nil
}

func (db *Database) CheckDuplicateUser(username string) error {
	_, err := db.RetriveUserIdFromCredentials(username, "")
	if err == jlog.ErrApiUserNoExist {
		return nil
	}
	return jlog.ErrApiMultipleUsers
}
