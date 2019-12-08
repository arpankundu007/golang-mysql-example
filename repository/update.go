package repository

import (
	"github.com/julienschmidt/httprouter"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/data"
	"mobile-specs-golang/utils"
	"net/http"
)

func UpdateData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	spec, err := utils.DecodeJSON(r)
	if err != nil {
		panic(err.Error())
	}else {
		if UpdateDataInDB(spec, id){
			utils.EncodeJSON(w, "Mobile updated successfully")
		}
	}
}

func UpdateDataInDB(mobile data.Mobile, id string) bool{
	db := GetDB()
	defer db.Close()
	update, err := db.Prepare("UPDATE " + constants.TableName + " SET brand=?, model=?, processor=?, ram=?, storage=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	_, err = update.Exec(mobile.Brand, mobile.Model, mobile.Processor, mobile.Ram, mobile.Storage, id)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}
