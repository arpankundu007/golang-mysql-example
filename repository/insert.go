package repository

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/data"
	"mobile-specs-golang/utils"
	"net/http"
)

func InsertData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := GetDB()
	defer db.Close()

	var spec data.Mobile
	err := json.NewDecoder(r.Body).Decode(&spec)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		id := utils.GetUUID()
		if id != "" {
			brand := spec.Brand
			model := spec.Model
			processor := spec.Processor
			ram := spec.Ram
			storage := spec.Storage
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
}
