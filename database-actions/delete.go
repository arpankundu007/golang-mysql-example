package database_actions

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"mobile-specs-golang/constants"
	"net/http"
)

func DeleteData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		db := GetDB(constants.DbName)
		defer db.Close()
		id := ps.ByName("id")

		del, err := db.Prepare("DELETE FROM " + constants.TableName + " WHERE ID=?")
		if err != nil {
			panic(err.Error())
		}
		_, _ = del.Exec(id)
		_, _ = io.WriteString(w, "Deleted successfully")
}
