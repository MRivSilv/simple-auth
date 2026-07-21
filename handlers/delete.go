package handlers

import ( 
	"encoding/json"
	"net/http"
	"simple-auth/store"
)

func DeleteUser (s *store.UserStore) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		userid, ok := r.Context().Value(userContextKey).(string)
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		var body struct{
			UserID string `json:"UserID"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		_, exists := s.Get(userid)
		if !exists {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		if !s.Delete(userid) {
			http.Error(w, "user not found", http.StatusNotFound)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}