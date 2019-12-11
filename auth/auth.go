package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/utils"
	"net/http"
	"time"
)

var signingKey = []byte(constants.AuthKey)

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "User"
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	generatedToken, err := token.SignedString(signingKey)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return generatedToken, err
}

func GenerateJWT() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := generateJWT()
		if err != nil {
			_, _ = io.WriteString(w, err.Error())
			return
		}
		utils.EncodeJSON(w, token)
	})
}

func IsAuthorized(next http.Handler) http.Handler {	//AuthMiddleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("error parsing jwt")
				}
				return signingKey, nil
			})

			if err != nil {
				_, _ = fmt.Fprintf(w, err.Error())
			}

			if token.Valid {	//TODO: Modify logic to authenticate users
				next.ServeHTTP(w, r)
			}
		} else {

			_, _ = fmt.Fprintf(w, "Not Authorized")
		}
	})
}
