package main

import (
	"github.com/julienschmidt/httprouter"
	"mobile-specs-golang/repository"
	"net/http"
)

func init() {
	db := repository.GetDB()
	defer db.Close()
	repository.CreateTableIfNotExists(db)
}

func main() {

	router := httprouter.New()

	router.GET("/mobile/mobiles/:id", repository.GetMobileInfo)

	router.GET("/mobile/all", repository.GetAllData)

	router.POST("/mobile/add", repository.InsertData)

	router.PUT("/mobile/update/:id", repository.UpdateData)

	router.GET("/drop", repository.DropTableIfExists)

	router.DELETE("/delete/{id}", repository.DeleteData)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}