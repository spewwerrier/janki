package db

import "time"

type Knob struct {
	Creation   time.Time
	KnobName   string
	ForkOf     int
	IsPublic   bool
	Identifier string
}

type KnobDescription struct {
	description string
	topics      []string
	todo        []string
	tor         []string
	refs        []string
	urls        []string
	ques        []string
	suggestions []string
	knob_id     int
}
