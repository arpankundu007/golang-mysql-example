package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"io/ioutil"
	"mobile-specs-golang/models"
	"mobile-specs-golang/utils"
	"net/http"
	"strconv"
	"time"
)

type CustomClaims struct {
	jwt.StandardClaims
	Scopes []string
}

func generateJWT(mins string, user models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				Audience:  "",
				ExpiresAt: getTokenExpiry(mins),
				IssuedAt:  time.Now().Unix(),
				Issuer:    "CAFU APP DMCC",
				NotBefore: 0,
				Subject:   user.Id,
			},
			Scopes: user.Scopes,
		})

	generatedToken, err := token.SignedString(signingKey)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return generatedToken, err
}

func getTokenExpiry(mins string) int64 {
	minsInt, _ := strconv.Atoi(mins)
	duration := time.Duration(minsInt)
	return time.Now().Add(time.Minute * duration).Unix()
}

func GetJWTToken() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user = models.User{}
		body, err := ioutil.ReadAll(r.Body)
		if err!=nil{
			utils.EncodeJSON(w, err.Error())
			return
		}else {
			err := json.Unmarshal(body, &user)
			if err != nil {
				utils.EncodeJSON(w, err.Error())
				return
			} else {
				mins := utils.GetParamFromRequestUrl(r, "exp")
				token, err := generateJWT(mins, user)
				if err != nil {
					_, _ = io.WriteString(w, err.Error())
					return
				}
				utils.EncodeJSON(w, token)
			}
		}
	})
}
