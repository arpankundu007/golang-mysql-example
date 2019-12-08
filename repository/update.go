package repository

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/data"
	"net/http"
)

func UpdateData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := GetDB()
	defer db.Close()
	id := ps.ByName("id")
	spec := data.Mobile{}
	err := json.NewDecoder(r.Body).Decode(&spec)
	if err != nil {
		panic(err.Error())
	}

	update, err := db.Prepare("UPDATE " + constants.TableName + " SET brand=?, model=?, processor=?, ram=?, storage=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	_, err = update.Exec(spec.Brand, spec.Model, spec.Processor, spec.Ram, spec.Storage, id)
	if err != nil {
		panic(err.Error())
	}
}
