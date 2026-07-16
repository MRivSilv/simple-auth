package handlers
import (
	"encoding/json"
	"net/http"
	"simple-auth/auth"
	"simple-auth/store"
)

type Credentials struct{
	Email string `json:"email"`
	Username string `json:"username"`
	Password string	`json:"password"`
}

func Register(s *store.UserStore) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		var creds Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "server error", http.StatusBadRequest)
			return
		}
		hash, err := auth.HashPassword(creds.Password)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		if !s.Create(creds.Username, hash){
			http.Error(w, "user exist", http.StatusConflict)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
