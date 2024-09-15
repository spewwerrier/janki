package db

type UsersDetails struct {
	Info        Info         `json:"Info"`
	Description Descriptions `json:"Descriptions"`
	Session     Session      `json:"Session"`
}

type Info struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

type Descriptions struct {
	Creation       int    `json:"Creation"`
	Image_url      string `json:"ImageUrl"`
	Description    string `json:"Description"`
	Existing_knobs int    `json:"ExistingKnobs"`
}

type Session struct {
	Cookie_string string `json:"CookieString"`
	Creation      int    `json:"Creation"`
	User_id       int    `json:"UserId"`
}
