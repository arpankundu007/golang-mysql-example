package main

import (
	"github.com/julienschmidt/httprouter"
	"mobile-specs-golang/auth"
	"mobile-specs-golang/repository"
	"net/http"
)

func init() {
	db := repository.GetDB()
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository.CreateTableIfNotExists(db)
}

func main() {

	router := httprouter.New()

	router.Handler(http.MethodGet, "/mobile/mobiles/", auth.IsAuthorized(repository.GetMobileInfo()))

	router.Handler(http.MethodGet, "/token", auth.GenerateJWT())

	router.Handler(http.MethodGet, "/mobile/all", auth.IsAuthorized(repository.GetAllData()))

	router.Handler(http.MethodPost, "/mobile/add", auth.IsAuthorized(repository.InsertData()))

	router.Handler(http.MethodPut, "/mobile/update/:id", auth.IsAuthorized(repository.UpdateData()))

	router.Handler(http.MethodDelete, "/delete/:id", auth.IsAuthorized(repository.DeleteData()))

	router.GET("/drop", repository.DropTableIfExists)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}