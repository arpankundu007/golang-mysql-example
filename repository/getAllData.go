package repository

import (
	"github.com/julienschmidt/httprouter"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/data"
	"mobile-specs-golang/utils"
	"net/http"
)

func GetAllData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := GetDB()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM " + constants.TableName + " ORDER BY id ASC ")

	if err != nil {
		panic(err.Error())
	}
	spec := data.Mobile{}
	var specAll []data.Mobile
	for selDB.Next() {
		var id, brand, model, processor, ram, storage string
		err := selDB.Scan(&id, &brand, &model, &processor, &ram, &storage)
		if err != nil {
			panic(err.Error())
		}
		spec.Id = id
		spec.Model = model
		spec.Brand = brand
		spec.Storage = storage
		spec.Ram = ram
		spec.Processor = processor
		specAll = append(specAll, spec)
	}
	utils.EncodeJSON(w, specAll)
}
