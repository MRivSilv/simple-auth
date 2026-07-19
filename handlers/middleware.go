package handlers

import (
	"net/http"
	"strings"
	"simple-auth/auth"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := auth.ParseToken(tokenStr)
		if err != nil || !token.Valid{
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}