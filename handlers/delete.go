package handlers

import ( 
	"encoding/json"
	"net/http"
	"simple-auth/auth"
	"simple-auth/store"
)

func DeleteUser (s *store.UserStore) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		username, ok := r.Context().Value(userContextKey).(string)
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		var body struct{
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		hash, exists := s.Get(username)
		if !exists || !auth.CheckPassword(hash, body.Password) {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		if !s.Delete(username) {
			http.Error(w, "user not found", http.StatusNotFound)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}