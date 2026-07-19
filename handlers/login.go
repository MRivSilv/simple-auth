package handlers

import (
	"encoding/json"
	"net/http"
	"simple-auth/auth"
	"simple-auth/store"
)

func Login (s *store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var creds Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
		}
		hash, ok := s.Get(creds.Username)
		if !ok || !auth.CheckPassword(hash, creds.Password){
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		token, err := auth.GenerateToken(creds.Username)
		if err != nil{
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}