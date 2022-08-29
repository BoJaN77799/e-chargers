package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// SECRET This should be stored as an environment variable
var SECRET = []byte("my_ultra_secret_key")

func ParseTokenStr(tokenStr string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("error while parsing jwt str")
		}
		return SECRET, nil
	})

	return token, err
}
