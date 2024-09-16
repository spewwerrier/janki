package api

import (
	"fmt"
	"janki/db"
	jankilog "janki/logs"
	"net/http"
)

type Users struct {
	Fields
	DB  *db.Database
	Log jankilog.JankiLog
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		u.Log.ErrorHttp(http.StatusBadRequest, "failed to parse form", w)
		return
	}

	// this is just a simple validation, move this to new file with all input validations and such
	username := r.Form.Get("username")
	if len(username) < 3 {
		u.Log.ErrorHttp(http.StatusBadRequest, "username should be more than 3 characters long", w)
		return
	}

	password := r.Form.Get("password")
	if len(password) < 8 {
		u.Log.ErrorHttp(http.StatusBadRequest, "password should be more than 8 characters long", w)
		return
	}

	session_key, err := u.DB.CreateNewUser(username, password, "https://example.com", "I am groot")
	if err != nil {
		u.Log.ErrorHttp(http.StatusBadRequest, "duplicate user", w)
		return
	}
	cookie := http.Cookie{
		Name:   "user_token",
		Value:  session_key,
		Path:   "/",
		MaxAge: 3600,
	}
	http.SetCookie(w, &cookie)
	_, _ = w.Write([]byte(cookie.Value))
}

func (u Users) CreateDescription(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		u.Log.ErrorHttp(http.StatusBadRequest, "failed to parse form", w)
		return
	}
	session_key := r.Form.Get("session_key")
	image_url := r.Form.Get("image_url")
	description := r.Form.Get("description")
	err = u.DB.CreateUserDescription(session_key, image_url, description)
	if err != nil {
		u.Log.ErrorHttp(http.StatusInternalServerError, "cannot create user description", w)
		return
	}
	_, _ = w.Write([]byte("created description"))
}

func (u Users) UpdateUserDescription(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		u.Log.ErrorHttp(http.StatusBadRequest, "failed to parse form", w)
		return
	}
	session_key := r.Form.Get("session_key")
	image_url := r.Form.Get("image_url")
	description := r.Form.Get("description")
	err = u.DB.UpdateUser(session_key, image_url, description)
	if err != nil {
		u.Log.ErrorHttp(http.StatusInternalServerError, "cannot update user description for some reason", w)
		return
	}
	_, _ = w.Write([]byte("updated description"))
}

func (u Users) Read(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		u.Log.ErrorHttp(http.StatusBadRequest, "failed to parse form", w)
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	session_key, err := u.DB.RetriveUserSession(username, password)
	if err != nil {
		u.Log.ErrorHttp(http.StatusInternalServerError, "cannot retrive user session", w)
		return
	}
	fmt.Println(session_key)
	u.Log.Println("gave session key to", username)
	_, _ = w.Write([]byte(session_key))

}

func (u Users) Error(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("error on user"))
}
