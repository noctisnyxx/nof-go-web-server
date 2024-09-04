package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"nof-go-web-server/internal/module"

	"github.com/julienschmidt/httprouter"
)

func NewSchedule(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	respBody := module.HttpBody{Status: http.StatusOK}
	data := new(module.ScheduleData)
	data.ScheduleId = data.GenerateScheduleId()
	body, err := io.ReadAll(request.Body)
	if err != nil {
		respBody.Status = http.StatusInternalServerError
		jsonRespBody, _ := json.Marshal(respBody)
		response.WriteHeader(respBody.Status)
		response.Write(jsonRespBody)
		return
	}
	json.Unmarshal(body, &data)
	respBody.Data = data
	jsonRespBody, err := json.Marshal(respBody)
	if err != nil {
		respBody.Status = http.StatusInternalServerError
		return
	}
	response.Write(jsonRespBody)

}

func EditSchedule(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}

func AbortSchedule(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}

func SearchSchedule(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}
