package repository

import (
	"github.com/julienschmidt/httprouter"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/data"
	"mobile-specs-golang/utils"
	"net/http"
)

func GetMobileInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := GetDB()
	defer db.Close()
	id := ps.ByName("id")
	selDB, err := db.Query("SELECT * FROM "+constants.TableName+" WHERE ID = ?", id)
	if err != nil {
		panic(err.Error())
	}
	spec := data.Mobile{}

	for selDB.Next() {
		var id, brand, model, processor, ram, storage string
		err := selDB.Scan(&id, &brand, &model, &processor, &ram, &storage)
		if err != nil {
			panic(err.Error())
		}
		spec.Id = id
		spec.Brand = brand
		spec.Model = model
		spec.Processor = processor
		spec.Ram = ram
		spec.Storage = storage
	}
	utils.EncodeJSON(w, spec)
}
