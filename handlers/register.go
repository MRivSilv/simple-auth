package handlers

import (
	"encoding/json"
	"net/http"
	"simple-auth/auth"
	"simple-auth/store"
)

type Credentials struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterUser(s *store.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "server error", http.StatusBadRequest)
			return
		}
		if creds.Email == "" || creds.Password == "" || creds.Username == "" {
			http.Error(w, "credential is missing", http.StatusBadRequest)
			return
		}
		hash, err := auth.HashPassword(creds.Password)
		if err != nil {
			http.Error(w, "error hashing password", http.StatusInternalServerError)
			return
		}

		_, err = s.CreateUser(creds.Email, creds.Username, hash)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
