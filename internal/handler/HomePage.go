package handler

import (
	"net/http"
	"nof-go-web-server/internal/utils"

	"github.com/julienschmidt/httprouter"
)

func ShowHomePage(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	res := utils.ResponseBody{
		Status: http.StatusOK,
		Data:   "welcome to home page!",
	}
	res.UpdateHttpResponse(writer, res.Status, res.Data)
}
