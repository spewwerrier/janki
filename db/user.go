package db

import "time"

type User struct {
	Name string `json:"Name"`
}

type UserDescription struct {
	Session        Session   `json:"Session"`
	Creation       time.Time `json:"Creation"`
	User           User      `json:"Info"`
	Image_url      string    `json:"ImageUrl"`
	Description    string    `json:"Description"`
	Existing_knobs int       `json:"ExistingKnobs"`
}

type Session struct {
	Creation time.Time `json:"Creation"`
	ApiKey   string    `json:"ApiKey"`
}
