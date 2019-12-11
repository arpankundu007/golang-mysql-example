package repository

import (
	"mobile-specs-golang/constants"
	"mobile-specs-golang/models"
	"mobile-specs-golang/utils"
	"net/http"
	"strings"
	"time"
)

func UpdateData() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r*http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/mobile/update/")
		spec, err := utils.DecodeJSON(r)
		if err != nil {
			panic(err.Error())
		}else {
			if UpdateDataInDB(spec, id){
				utils.EncodeJSON(w, "Mobile updated successfully")
			}
		}
	})
}

func UpdateDataInDB(mobile models.Mobile, id string) bool{
	db := GetDB()
	defer db.Close()
	update, err := db.Prepare("UPDATE " + constants.TableName + " SET brand=?, model=?, processor=?, ram=?, storage=?, updatedAt=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	_, err = update.Exec(mobile.Brand, mobile.Model, mobile.Processor, mobile.Ram, mobile.Storage, time.Now().Unix(), id)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}
