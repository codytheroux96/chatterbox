package main

import (
	"net/http"

	"github.com/codytheroux96/chatterbox/internal"
)

func main() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if err := internal.RegisterUser(r.FormValue("username"), r.FormValue("password")); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		token, err := internal.LoginUser(r.FormValue("username"), r.FormValue("password"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write([]byte(token))
	})

	http.ListenAndServe(":8080", nil)
}