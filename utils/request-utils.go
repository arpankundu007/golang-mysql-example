package utils

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetParamFromRequestUrl(r *http.Request, param string) string{
	return httprouter.ParamsFromContext(r.Context()).ByName(param)
}
