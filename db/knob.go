package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

var KnobItemTable = []string{"knobquestions", "knobreferences", "knobsuggestions", "knobthingstoread", "knobtodo", "knobtopics", "knoburls"}

type Knob struct {
	Creation    time.Time `json:"Creation"`
	KnobItems   KnobItems `json:"KnobItems"`
	KnobName    string    `json:"KnobName"`
	Identifier  string    `json:"Identifier"`
	ForkOf      string    `json:"ForkOf"`
	Description string    `json:"Description"`
	IsPublic    bool      `json:"IsPublic"`
}

type KnobReferences struct {
	References []string
}

type KnobTopics struct {
	Topics []string
}

type KnobThingsToRead struct {
	ThingsToRead []string
}

type KnobUrls struct {
	URLS []string
}

type KnobQuestions struct {
	Questions []string
}

type KnobSuggestions struct {
	Suggestions []string
}

type KnobTodo struct {
	Todo []string
}

type KnobItems struct {
	KnobReferences   `json:"KnobReferences"`
	KnobTopics       `json:"KnobTopics"`
	KnobThingsToRead `json:"KnobThingsToRead"`
	KnobUrls         `json:"KnobUrls"`
	KnobQuestions    `json:"KnobQuestions"`
	KnobSuggestions  `json:"KnobSuggestions"`
	KnobTodo         `json:"KnobTodo"`
}

// TODO: DELETE THIS
type KnobDescription struct {
	Knob        Knob                 `json:"Knob"`
	Description string               `json:"Description"`
	Topics      pgtype.Array[string] `json:"Topics"`
	Todo        pgtype.Array[string] `json:"Todo"`
	Tor         pgtype.Array[string] `json:"Tor"`
	Refs        pgtype.Array[string] `json:"Refs"`
	Urls        pgtype.Array[string] `json:"Urls"`
	Ques        pgtype.Array[string] `json:"Ques"`
	Suggestions pgtype.Array[string] `json:"Suggestions"`
}
