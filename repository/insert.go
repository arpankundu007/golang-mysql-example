package repository

import (
	"mobile-specs-golang/constants"
	"mobile-specs-golang/models"
	"mobile-specs-golang/utils"
	"net/http"
	"time"
)

func InsertData() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mobile, err := utils.DecodeJSON(r)	//Decode the request body and format it as per "spec"
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			InsertIntoDb(mobile)
		}
	})
}

func InsertIntoDb(mobile models.Mobile){
	db := GetDB()
	defer db.Close()
	id, err := utils.GetUUID()
	if err == nil {
		brand := mobile.Brand
		model := mobile.Model
		processor := mobile.Processor
		ram := mobile.Ram
		storage := mobile.Storage
		createdAt := time.Now().Unix()
		updatedAt := time.Now().Unix()
		stmt, err := db.Prepare("INSERT INTO " + constants.TableName + " (id, brand, model, processor, ram, storage, createdAt, updatedAt) VALUES (?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		_, err = stmt.Exec(id, brand, model, processor, ram, storage, createdAt, updatedAt)
		if err != nil {
			panic(err.Error())
		}
	} else {
		panic("UUID generation failed")
	}
}
