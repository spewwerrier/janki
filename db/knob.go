package db

import (
	"time"

	"github.com/lib/pq"
)

type Knob struct {
	Creation   time.Time
	KnobName   string
	Identifier string
	ForkOf     int
	IsPublic   bool
}

type KnobDescription struct {
	Description string         `json:"Description"`
	Topics      pq.StringArray `json:"Topics"`
	Todo        pq.StringArray `json:"Todo"`
	Tor         pq.StringArray `json:"Tor"`
	Refs        pq.StringArray `json:"Refs"`
	Urls        pq.StringArray `json:"Urls"`
	Ques        pq.StringArray `json:"Ques"`
	Suggestions pq.StringArray `json:"Suggestions"`
}
