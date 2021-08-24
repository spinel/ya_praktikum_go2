package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name     string
	LastName string
}

func UserViewHandler(users map[string]User) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userId := r.FormValue("user_id")

		if r.FormValue("user_id") == "" {
			http.Error(rw, "userId is empty", http.StatusBadRequest)
		}

		user, ok := users[userId]
		//
		if !ok {
			http.Error(rw, "user not found", http.StatusNotFound)
		}

		jsonUser, err := json.Marshal(user)
		if err != nil {
			http.Error(rw, "cant provide a json. internal error", http.StatusBadGateway)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_, err = rw.Write(jsonUser)
		if err != nil {
			http.Error(rw, "cant provide a json. internal error", http.StatusBadGateway)
		}
	}
}
