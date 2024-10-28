package db

import (
	"errors"
	"fmt"

	"janki/jlog"
)

const (
	RetrieveTopics       = "select topics from knobtopics where knob_id=$1"
	RetrieveQuestions    = "select questions from knobquestions where knob_id=$1"
	RetrieveReferences   = "select refs from knobreferences where knob_id=$1"
	RetrieveSuggestions  = "select suggestions from knobsuggestions where knob_id=$1"
	RetrieveThingsToRead = "select thingstoread from knobthingstoread where knob_id=$1"
	RetrieveTodo         = "select todo from  knobtodo where knob_id=$1"
	RetrieveUrls         = "select urls from knoburls where knob_id=$1"
)

// fills up the knob with its every detail using its identifier
func (db *Database) RetrieveKnobItem(knob *Knob) {
	knob_id, err := db.GetKnobIdFromIdentifier(knob.Identifier)
	if err != nil {
		db.log.Error("RetrieveKnobItem failed to retrieve identifir")
		return
	}
	{
		query := "select knob_name, description, creation, ispublic from knobs where id=$1"
		row := db.QueryRow(query, knob_id)
		row.Scan(&knob.KnobName, &knob.Description, &knob.Creation, &knob.IsPublic)
	}

	// for topics
	{
		rows, err := db.Query(RetrieveTopics, knob_id)
		if err != nil {
			db.log.Error("RetrieveKnobItem failed to retrieve knob")
			return
		}

		for rows.Next() {
			var row string
			rows.Scan(&row)
			knob.KnobItems.Topics = append(knob.KnobItems.Topics, row)
		}
	}

	// for questions
	{
		rows, err := db.Query(RetrieveQuestions, knob_id)
		if err != nil {
			db.log.Error("RetrieveKnobItem failed to retrieve questions")
			return
		}

		for rows.Next() {
			var row string
			rows.Scan(&row)
			knob.KnobItems.Questions = append(knob.KnobItems.Questions, row)
		}
	}
	// for References
	{
		rows, err := db.Query(RetrieveReferences, knob_id)
		if err != nil {
			db.log.Error("RetrieveKnobItem failed to retrieve References")
			return
		}

		for rows.Next() {
			var row string
			rows.Scan(&row)
			knob.KnobItems.References = append(knob.KnobItems.References, row)
		}
	}
	// for Suggestions
	{
		rows, err := db.Query(RetrieveSuggestions, knob_id)
		if err != nil {
			db.log.Error("RetrieveKnobItem failed to retrieve Suggestions")
			return
		}

		for rows.Next() {
			var row string
			rows.Scan(&row)
			knob.KnobItems.Suggestions = append(knob.KnobItems.Suggestions, row)
		}
	}
	// for ThingsToRead
	{
		rows, err := db.Query(RetrieveThingsToRead, knob_id)
		if err != nil {
			db.log.Error("RetrieveKnobItem failed to retrieve ThingsToRead")
			return
		}

		for rows.Next() {
			var row string
			rows.Scan(&row)
			knob.KnobItems.ThingsToRead = append(knob.KnobItems.ThingsToRead, row)
		}
	}
	// for Todo
	{
		rows, err := db.Query(RetrieveTodo, knob_id)
		if err != nil {
			db.log.Error("RetrieveKnobItem failed to retrieve Todo")
			return
		}

		for rows.Next() {
			var row string
			rows.Scan(&row)
			knob.KnobItems.Todo = append(knob.KnobItems.Todo, row)
		}
	}
	// for Urls
	{
		rows, err := db.Query(RetrieveUrls, knob_id)
		if err != nil {
			db.log.Error("RetrieveKnobItem failed to retrieve Urls")
			return
		}

		for rows.Next() {
			var row string
			rows.Scan(&row)
			knob.KnobItems.URLS = append(knob.KnobItems.URLS, row)
		}
	}
}

// returns every knob that a user has. It only returns a brief overall unlike RetrieveKnobItem
func (db *Database) RetrieveUserKnobs(api_key string) ([]Knob, error) {
	id, err := db.RetriveUserIdFromApi(api_key)
	if err != nil {
		db.log.Error("GetUserKnob failed to get user id")
		return nil, err
	}
	query := "select knob_name, creation,ispublic, identifier, description from knobs  where user_id = $1 order by creation desc"
	result, err := db.Query(query, id)
	if err != nil {
		db.log.Error("GetUserKnobs " + err.Error())
		return nil, err
	}
	var i int
	var knob Knob
	var knobs []Knob
	fmt.Println(query, id)
	for result.Next() {
		i++
		result.Scan(&knob.KnobName, &knob.Creation, &knob.IsPublic, &knob.Identifier, &knob.Description)
		knobs = append(knobs, knob)
	}
	if i < 1 {
		db.log.Warning("GetUserKnobs no knob exists")
		return nil, jlog.ErrNoKnobExists
	}

	return knobs, nil
}

// returns ErrNoKnobExists indicating new knob can be made == 0
// returns ErrKnobAlreadyExists indicating new knob cannot be made (we should never hit this) > 1
// return knobId if there is already a knob == 1
func (db *Database) GetKnobId(api_key string, knob_name string) (int, error) {
	id, err := db.GetUserId(api_key)
	if err != nil {
		return -1, err
	}
	query := "select id from knobs where user_id = $1 and knob_name = $2"
	result, err := db.Query(query, id, knob_name)
	if err != nil {
		return -1, err
	}
	var i int
	var knobId int
	for result.Next() {
		i++
		result.Scan(&knobId)
	}
	if i == 0 {
		return -1, jlog.ErrNoKnobExists
	}
	if i > 1 {
		return -1, jlog.ErrKnobAlreadyExists
	}
	return knobId, nil
}

func (db *Database) GetKnobIdFromIdentifier(identifier string) (int, error) {
	query := "select id from knobs where identifier = $1"
	result, err := db.Query(query, identifier)
	if err != nil {
		return -1, err
	}
	var i int
	var knobId int
	for result.Next() {
		i++
		result.Scan(&knobId)
	}
	if i == 0 {
		return -1, jlog.ErrNoKnobExists
	}
	if i > 1 {
		return -1, jlog.ErrKnobAlreadyExists
	}
	return knobId, nil
}

func (db *Database) GetUserIdFromKnobIdentifier(identifier string) (int, error) {
	query := "select user_id from knobs where identifier=$1"
	result := db.QueryRow(query, identifier)
	var id int
	result.Scan(&id)
	return id, nil
}

// verifies wether the user is editing their's knob or not
func (db *Database) AuthorizeUserKnob(api string, identifier string) (bool, error) {
	fmt.Println(api, identifier)

	userIdApi, err := db.GetUserId(api)
	if err != nil {
		db.log.Error("GetKnobDescriptions failed to get user id from API")
		return false, err
	}
	userIdIdentifier, err := db.GetUserIdFromKnobIdentifier(identifier)
	if err != nil {
		db.log.Error("GetKnobDescriptions failed to get user id from identifier")
		return false, err
	}

	if userIdApi == userIdIdentifier {
		return true, nil
	} else {
		return false, nil
	}
}

// if api key is given and its correct then it returns all public and private knobs else get only public knob
func (db *Database) RetrieveKnobDescriptions(api string, identifier string) (Knob, error) {
	knob := Knob{}
	knob.Identifier = identifier

	knob_id, _ := db.GetKnobIdFromIdentifier(identifier)

	query := "select knobs.ispublic from knobs where id=$1"
	rows := db.QueryRow(query, knob_id)
	var d bool
	rows.Scan(&d)

	knob.IsPublic = d

	isAuthorized, _ := db.AuthorizeUserKnob(api, identifier)
	if !isAuthorized && !knob.IsPublic {
		fmt.Println("unauthorized access")
		return knob, errors.New("unauthorized to view")
	}

	db.RetrieveKnobItem(&knob)

	return knob, nil
}
