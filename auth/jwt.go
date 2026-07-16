package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("change-it-later-it-goes-on-env-vars-top-secret-ass-key")

func GenerateToken(username string) (string, error){
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(15* time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error){
		return jwtKey, nil
	})
}