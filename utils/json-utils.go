package utils

import (
	"encoding/json"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/data"
	"net/http"
)

func EncodeJSON(w http.ResponseWriter, data interface{}){
	w.Header().Set(constants.ContentTypeHeader, constants.ContentTypeValue)
	parseError := json.NewEncoder(w).Encode(&data)
	if parseError != nil {
		panic(parseError.Error())
	}
}

func DecodeJSON(r *http.Request) (data.Mobile, error){
	var mobile data.Mobile
	err := json.NewDecoder(r.Body).Decode(&mobile)
	return mobile, err
}
