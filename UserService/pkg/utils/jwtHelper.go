package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strings"
)

func ExtractAccessTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("request doesn't contain header 'Authorization'")
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	return splitToken[1], nil
}

func parseTokenAndGet(r *http.Request, key string) (string, error) {
	tokenString, err := ExtractAccessTokenFromHeader(r)
	if err != nil {
		return "", err
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	claims := token.Claims.(jwt.MapClaims)
	return claims[key].(string), nil
}

func GetUserIDFromToken(r *http.Request) (uuid.UUID, error) {
	id, err := parseTokenAndGet(r, "id")
	if err != nil {
		return uuid.UUID{}, err
	}

	userId, _ := uuid.FromString(id)
	return userId, nil
}
