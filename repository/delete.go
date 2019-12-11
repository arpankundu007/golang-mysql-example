package repository

import (
	"io"
	"mobile-specs-golang/constants"
	"net/http"
	"strings"
)

func DeleteData() http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := GetDB()
		defer db.Close()
		id := strings.TrimPrefix(r.URL.Path, "/mobile/update/")

		del, err := db.Prepare("DELETE FROM " + constants.TableName + " WHERE ID=?")
		if err != nil {
			panic(err.Error())
		}
		_, _ = del.Exec(id)
		_, _ = io.WriteString(w, "Deleted successfully")
	})

}
