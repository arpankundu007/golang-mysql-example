package repository

import (
	"fmt"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/models"
	"mobile-specs-golang/utils"
	"net/http"
	"strings"
)

func GetMobileInfo() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := GetDB()
		defer db.Close()
		id := strings.TrimPrefix(r.URL.Path, "/mobile/mobiles/")
		fmt.Println(id)
		selDB, err := db.Query("SELECT * FROM "+constants.TableName+" WHERE ID = ?", id)
		if err != nil {
			panic(err.Error())
		}
		spec := models.Mobile{}

		for selDB.Next() {
			var id, brand, model, processor, ram, storage, createdAt, updatedAt string
			err := selDB.Scan(&id, &brand, &model, &processor, &ram, &storage, &createdAt, &updatedAt)
			if err != nil {
				panic(err.Error())
			}
			spec.Id = id
			spec.Brand = brand
			spec.Model = model
			spec.Processor = processor
			spec.Ram = ram
			spec.Storage = storage
			spec.CreatedAt = createdAt
			spec.UpdatedAt = updatedAt
		}
		utils.EncodeJSON(w, spec)
	})

}
