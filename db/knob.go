package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Knob struct {
	Creation   time.Time
	KnobName   string
	Identifier string
	ForkOf     string
	IsPublic   bool
}

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
