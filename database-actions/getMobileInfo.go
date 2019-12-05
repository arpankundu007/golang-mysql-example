package database_actions

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/data"
	"net/http"
)

func GetMobileInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := GetDB(constants.DbName)
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	parseError := json.NewEncoder(w).Encode(&spec)

	if parseError != nil {
		panic(parseError.Error())
	}

}
