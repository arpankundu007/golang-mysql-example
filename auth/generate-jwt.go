package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"mobile-specs-golang/models"
	"mobile-specs-golang/utils"
	"net/http"
	"strconv"
	"time"
)


func generateJWT(mins string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	setTokenPayload(claims)
	setTokenExpiry(mins, claims)

	generatedToken, err := token.SignedString(signingKey)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return generatedToken, err
}

func setTokenPayload(claims jwt.MapClaims){
	claims["user"] = utils.StructToJSONString(models.User{
		Id:    "user_id",
		Name:  "James Bond",
		Phone: "007",
	})
}

func setTokenExpiry(mins string, claims jwt.MapClaims){
	minsInt, _ := strconv.Atoi(mins)
	duration := time.Duration(minsInt)
	claims["exp"] = time.Now().Add(time.Minute * duration).Unix()
}

func GetJWTToken() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mins := utils.GetParamFromRequestUrl(r, "exp")
		token, err := generateJWT(mins)
		if err != nil {
			_, _ = io.WriteString(w, err.Error())
			return
		}
		utils.EncodeJSON(w, token)
	})
}
