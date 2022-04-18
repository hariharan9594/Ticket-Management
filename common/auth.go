package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Generate JWT token
func GenerateJWT(id uint, isadm bool) (string, error) {
	tokenContent := jwt.MapClaims{
		"UserInfo":   struct {
			Id      uint
			IsAdmin bool
		}{id, isadm},
		"exp": time.Now().Add(time.Minute * 72).Unix(),
	}

	t := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	tokenString, err := t.SignedString([]byte("TokenPassword"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}