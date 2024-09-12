package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResponseBody struct {
	Status int
	Data   interface{}
}

func (res ResponseBody) UpdateHttpResponse(writer http.ResponseWriter, newStatus int, newData interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	body := ResponseBody{
		Status: newStatus,
		Data:   newData,
	}
	byteBody, err := json.Marshal(body)
	if err != nil {
		http.Error(writer, "Failed to encode to json file", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(body.Status)
	writer.Write(byteBody)
}

func HttpRequestBodyReader[T interface{}](request *http.Request) (T, error) {
	var data T
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return data, fmt.Errorf("an error ocured: %s", err.Error())
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return data, fmt.Errorf("an error ocured: %s", err.Error())
	}
	return data, nil
}
