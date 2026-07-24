package handlers

import (
	"encoding/json"
	"net/http"
	"simple-auth/store"
)

type AppCredentials struct {
	AppName     string `json:"appname"`
	Description string `json:"description"`
}

func CreateApp(s *store.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds AppCredentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "server error", http.StatusBadRequest)
			return
		}
		if creds.AppName == "" || creds.Description == "" {
			http.Error(w, "credentials missing", http.StatusBadRequest)
			return
		}

		_, err := s.CreateApp(creds.AppName, creds.Description)
		if err != nil {
			http.Error(w, "Error creating app", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
//TODO: DELETE APP

//TODO: EDIT APP
