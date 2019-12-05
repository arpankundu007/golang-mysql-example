package database_actions

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"io"
	"mobile-specs-golang/constants"
	"net/http"
)

func DropTableIfExists(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := sql.Open("mysql", "root:1q2w3e4r@tcp(127.0.0.1:3306)/"+constants.TableName)
	if err != nil {
		panic(err.Error())
	} else {
		_, err = db.Exec("DROP TABLE IF EXISTS " + constants.TableName + ";")
		if err != nil {
			panic(err.Error())
		} else {
			fmt.Println("Table dropped successfully")
			_, _ = io.WriteString(w, "Table dropped successfully")
		}
	}
}
