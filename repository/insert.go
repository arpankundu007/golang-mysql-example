package repository

import (
	"github.com/julienschmidt/httprouter"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/data"
	"mobile-specs-golang/utils"
	"net/http"
)

func InsertData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mobile, err := utils.DecodeJSON(r)	//Decode the request body and format it as per "spec"
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		InsertIntoDb(mobile)
	}
}

func InsertIntoDb(mobile data.Mobile){
	db := GetDB()
	defer db.Close()
	id, err := utils.GetUUID()
	if err == nil {
		brand := mobile.Brand
		model := mobile.Model
		processor := mobile.Processor
		ram := mobile.Ram
		storage := mobile.Storage
		stmt, err := db.Prepare("INSERT INTO " + constants.TableName + " (id, brand, model, processor, ram, storage) VALUES (?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		_, err = stmt.Exec(id, brand, model, processor, ram, storage)
		if err != nil {
			panic(err.Error())
		}
	} else {
		panic("UUID generation failed")
	}
}
