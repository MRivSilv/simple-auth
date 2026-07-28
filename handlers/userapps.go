package handlers

import (
	"encoding/json"
	"net/http"
	"simple-auth/store"
)

type UserApps struct {
	AppID  string `json:"appid"`
	UserID string `json:"userid"`
}

func AddToApp(s *store.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userapps UserApps
		if err := json.NewDecoder(r.Body).Decode(&userapps); err != nil {
			http.Error(w, "server error", http.StatusBadRequest)
			return
		}
		if userapps.AppID == "" || userapps.UserID == "" {
			http.Error(w, "ID Missing", http.StatusBadRequest)
			return
		}
		err := s.AddUserToApp(userapps.UserID, userapps.AppID)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func RemoveFromApp(s *store.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userapps UserApps
		if err := json.NewDecoder(r.Body).Decode(&userapps); err != nil {
			http.Error(w, "server error", http.StatusBadRequest)
			return
		}
		if userapps.AppID == "" || userapps.UserID == "" {
			http.Error(w, "IDs missing", http.StatusBadRequest)
			return
		}
		err := s.RemoveUserFromApp(userapps.UserID, userapps.AppID)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
