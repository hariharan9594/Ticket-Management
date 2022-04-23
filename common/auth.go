package common

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Generate JWT token
func GenerateJWT(id uint, isadm bool, uname string) (string, error) {
	tokenContent := jwt.MapClaims{
		"UserInfo": struct {
			Id       uint
			IsAdmin  bool
			UserName string
		}{id, isadm, uname},
		"exp": time.Now().Add(time.Minute * 72).Unix(),
	}
	fmt.Println("userInfo: ", tokenContent["UserInfo"])

	t := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	tokenString, err := t.SignedString([]byte("TokenPassword"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
