package module

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpBody struct {
	Status int
	Data   interface{}
}

func (res HttpBody) UpdateHttpResponse(response http.ResponseWriter, newStatus int, newData interface{}) {
	body := HttpBody{
		Status: newStatus,
		Data:   newData,
	}
	byteBody, err := json.Marshal(body)
	if err != nil {
		http.Error(response, "Failed to encode to json file", http.StatusInternalServerError)
		return
	}
	response.Write(byteBody)
}

func HttpRequestBodyJsonReader(request *http.Request) (data any, err error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("an error ocured: %s", err.Error())
	}
	var structData struct{}
	if err = json.Unmarshal(body, &structData); err != nil {
		return nil, fmt.Errorf("an error ocured: %s", err.Error())
	}
	return structData, nil
}
