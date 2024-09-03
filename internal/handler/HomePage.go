package handler

import (
	"encoding/json"
	"net/http"
	"nof-go-web-server/internal/structs"

	"github.com/julienschmidt/httprouter"
)

func ShowHomePage(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	bodyMessage := structs.HttpResp{
		Status: http.StatusText(http.StatusOK),
		Data:   "Welcome to homepage",
	}
	jsonBodyMessage, _ := json.Marshal(bodyMessage)
	response.Write(jsonBodyMessage)
}
