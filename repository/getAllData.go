package repository

import (
	"mobile-specs-golang/constants"
	"mobile-specs-golang/models"
	"mobile-specs-golang/utils"
	"net/http"
)

func GetAllData() http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		db := GetDB()
		defer db.Close()

		selDB, err := db.Query("SELECT * FROM " + constants.TableName + " ORDER BY id ASC ")

		if err != nil {
			panic(err.Error())
		}
		mobile := models.Mobile{}
		var mobiles []models.Mobile
		for selDB.Next() {
			var id, brand, model, processor, ram, storage, createdAt, updatedAt string
			err := selDB.Scan(&id, &brand, &model, &processor, &ram, &storage, &createdAt, &updatedAt)
			if err != nil {
				panic(err.Error())
			}
			mobile.Id = id
			mobile.Model = model
			mobile.Brand = brand
			mobile.Storage = storage
			mobile.Ram = ram
			mobile.Processor = processor
			mobile.CreatedAt = createdAt
			mobile.UpdatedAt = updatedAt
			mobiles = append(mobiles, mobile)
		}
		utils.EncodeJSON(w, mobiles)
	})

}
