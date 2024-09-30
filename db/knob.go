package db

type Knob struct {
	Id       int
	UsersId  int
	KnobName string
	Creation int
	ForkOf   int
	IsPublic bool
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
