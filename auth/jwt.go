package auth

import (
	"crypto/rsa"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {
	if pkPath := os.Getenv("JWT_PRIVATE_KEY"); pkPath != "" {
		data, err := os.ReadFile(pkPath)
		if err != nil {
			panic("failed to read JWT_PRIVATE_KEY: " + err.Error())
		}
		privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(data)
		if err != nil {
			panic("failed to parse JWT_PRIVATE_KEY: " + err.Error())
		}
	}

	pkPath := os.Getenv("JWT_PUBLIC_KEY")
	if pkPath == "" {
		pkPath = "keys/public.pem"
	}
	data, err := os.ReadFile(pkPath)
	if err != nil {
		panic("failed to read JWT_PUBLIC_KEY: " + err.Error())
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(data)
	if err != nil {
		panic("failed to parse JWT_PUBLIC_KEY: " + err.Error())
	}
}

func GenerateToken(username string) (string, error) {
	if privateKey == nil {
		panic("JWT_PRIVATE_KEY not configured")
	}

	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

func ParseToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
}
