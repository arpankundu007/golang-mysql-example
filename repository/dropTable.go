package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"io"
	"mobile-specs-golang/constants"
	"net/http"
)

func DropTableIfExists(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := GetDB()
	_, err := db.Exec("DROP TABLE IF EXISTS " + constants.TableName + ";")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Table dropped successfully. ")
		_, _ = io.WriteString(w, "Table dropped successfully. Restart server for changes")
	}
}
