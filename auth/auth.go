package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/utils"
	"net/http"
	"strings"
)

var signingKey = []byte(constants.AuthKey)

func IsAuthorized(next http.Handler) http.Handler { //AuthMiddleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken, err := getTokenFromHeader(r)
		if err == nil {
			token, err := jwt.ParseWithClaims(headerToken, &CustomClaims{}, verifyFunc) //Verifies the authenticity of accessToken
			if err != nil {
				_, _ = fmt.Fprintf(w, err.Error())
			} else {
				/* --------------------------------------------------------------------------*/
				//TODO: Modify logic to authenticate users
				if token.Valid {
					claims := token.Claims.(*CustomClaims)
					if utils.Contains(getScopes(claims), "mobiles") {
						next.ServeHTTP(w, r)
					} else {
						utils.EncodeJSON(w, "Client does not have sufficient scopes")
					}
				}

				/* --------------------------------------------------------------------------*/
			}
		} else {
			_, _ = fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func getScopes(claims *CustomClaims) []string {
	return claims.Scopes
}

func verifyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("error parsing jwt")
	}
	return signingKey, nil
}

func getTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	splitAuthHeader := strings.Split(authHeader, " ")
	if len(splitAuthHeader) != 2 {
		return "", errors.New("invalid auth header")
	}
	return splitAuthHeader[1], nil
}
