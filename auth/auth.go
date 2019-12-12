package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"mobile-specs-golang/constants"
	"net/http"
	"strings"
)

var signingKey = []byte(constants.AuthKey)

func IsAuthorized(next http.Handler) http.Handler { //AuthMiddleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken, err := getTokenFromHeader(r)
		if err == nil {
			token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("error parsing jwt")
				}
				return signingKey, nil
			})

			if err != nil {
				_, _ = fmt.Fprintf(w, err.Error())
			}

			if token.Valid { //TODO: Modify logic to authenticate users
				next.ServeHTTP(w, r)
			}
		} else {
			_, _ = fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func getTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	splitAuthHeader := strings.Split(authHeader, " ")
	if len(splitAuthHeader) != 2 {
		return "", errors.New("invalid auth header")
	}
	return splitAuthHeader[1], nil
}
