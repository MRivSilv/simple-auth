package handlers

import (
	"net/http"
	"simple-auth/store"
)

func DeleteUser(s *store.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid, ok := r.Context().Value(userContextKey).(string)
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		_, err := s.GetUser(userid)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		_, err = s.DeleteUser(userid)
		if err != nil {
			http.Error(w, "error deleting user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
