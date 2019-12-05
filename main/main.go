package main

import (
	"github.com/julienschmidt/httprouter"
	"mobile-specs-golang/constants"
	"mobile-specs-golang/database-actions"
	"net/http"
)

func init() {
	db := database_actions.GetDB(constants.DbName)
	defer db.Close()
	database_actions.CreateTableIfNotExists(db)
}

func main() {

	router := httprouter.New()

	router.GET("/mobile/mobiles/:id", database_actions.GetMobileInfo)

	router.GET("/mobile/all", database_actions.GetAllData)

	router.POST("/mobile/add", database_actions.InsertData)

	router.PUT("/mobile/update/:id", database_actions.UpdateData)

	router.GET("/drop", database_actions.DropTableIfExists)

	router.DELETE("/delete/{id}", database_actions.DeleteData)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}