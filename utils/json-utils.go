package utils

import (
	"encoding/json"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/models"
	"net/http"
)

func StructToJSONString(data interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err!=nil{
		panic(err.Error())
		return ""
	}
	return string(jsonBytes)
}

func EncodeJSON(w http.ResponseWriter, data interface{}){
	w.Header().Set(constants.ContentTypeHeader, constants.ContentTypeValue)
	parseError := json.NewEncoder(w).Encode(&data)
	if parseError != nil {
		panic(parseError.Error())
	}
}

func DecodeJSON(r *http.Request) (models.Mobile, error){
	var data models.Mobile
	err := json.NewDecoder(r.Body).Decode(&data)
	return data, err
}
