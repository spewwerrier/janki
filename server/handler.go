package server

import (
	"net/http"

	"janki/api"
)

func Middleware(handler http.Handler) http.Handler {
	// currently only sets up cors. Plans to add a debug
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		handler.ServeHTTP(w, r)
	})
}

func Handler(api *api.Api) *http.ServeMux {
	h := http.NewServeMux()

	h.HandleFunc("/users/create", api.Users.Create)
	h.HandleFunc("/users/update", api.Users.UpdateUserDescription)
	h.HandleFunc("/users/read", api.Users.Read)
	h.HandleFunc("/users/", api.Users.Error)

	h.HandleFunc("/knob/create", api.Knob.Create)
	h.HandleFunc("/knob/read", api.Knob.Read)
	h.HandleFunc("/knob/update", api.Knob.Update)

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Your error knows no bounds"))
	})

	return h
}
