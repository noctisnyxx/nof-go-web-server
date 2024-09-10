package handler

import (
	"net/http"
	"nof-go-web-server/internal/module"

	"github.com/julienschmidt/httprouter"
)

func ShowHomePage(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	module.UpdateHttpResponse(response, 200, "Welcome to homepage!")
}
