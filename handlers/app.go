package handlers

import (
	"encoding/json"
	"net/http"
	"simple-auth/store"
)

type AppCredentials struct {
	AppID       string `json:"appid"`
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

func RemoveApp(s *store.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds AppCredentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "server error", http.StatusBadRequest)
			return
		}
		if creds.AppID == "" {
			http.Error(w, "AppID missing", http.StatusBadRequest)
			return
		}

		ok, err := s.DeleteApp(creds.AppID)
		if err != nil {
			http.Error(w, "Error deleting app", http.StatusInternalServerError)
			return
		}
		if !ok {
			http.Error(w, "App not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

//TODO: EDIT APP
