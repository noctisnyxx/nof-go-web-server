package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nof-go-web-server/internal/structs"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Item struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"picture_link"`
}

func AddItem(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	itemList := []Item{}
	response.Header().Set("Content-Type", "application/json")
	storage_path := "./assets/item_list.json"
	_, errNotExist := os.Stat(storage_path)

	if os.IsNotExist(errNotExist) {
		file, err := os.Create(storage_path)
		if err != nil {
			jsonrespBody, _ := json.Marshal(structs.HttpResp{
				Status: strconv.Itoa(http.StatusInternalServerError) + ":" + http.StatusText(http.StatusInternalServerError),
				Data:   nil,
			})
			response.Write(jsonrespBody)
			return
		}
		defer file.Close()
		jsonrespBody, _ := json.Marshal(structs.HttpResp{
			Status: strconv.Itoa(http.StatusOK) + ":" + http.StatusText(http.StatusOK),
			Data:   nil,
		})
		response.Write(jsonrespBody)
	}

	byteData, err := os.ReadFile(storage_path)
	if err != nil {
		fmt.Println(err.Error())
		response.WriteHeader(http.StatusInternalServerError) //
		return
	}
	price, _ := strconv.ParseFloat(request.FormValue("Price"), 64)
	item := Item{
		Id:          request.FormValue("Id"),
		Name:        request.FormValue("Name"),
		Category:    request.FormValue("Category"),
		Price:       price,
		Description: request.FormValue("Description"),
		Image:       request.FormValue("Image"),
	}

	json.Unmarshal(byteData, &itemList)
	for _, currentData := range itemList {
		if item.Id == currentData.Id {
			response.WriteHeader(http.StatusBadRequest)
			jsonrespBody, _ := json.Marshal(structs.HttpResp{
				Status: strconv.Itoa(http.StatusOK) + ": The same ID has founded",
				Data:   item.Id,
			})
			response.Write(jsonrespBody)
			return
		}
	}

	file, err := os.OpenFile(storage_path, os.O_RDWR, 0644)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		jsonrespBody, _ := json.Marshal(structs.HttpResp{
			Status: strconv.Itoa(http.StatusBadRequest) + ":" + http.StatusText(http.StatusBadRequest),
			Data:   nil,
		})
		response.Write(jsonrespBody)
		return
	}

	itemList = append(itemList, item)
	data, _ := json.MarshalIndent(itemList, " ", " ")
	file.Write(data)
	response.WriteHeader(http.StatusOK)
	jsonrespBody, _ := json.Marshal(item)
	response.Write(jsonrespBody)
}

func UpdateItem(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	storage_path := "./assets/item_list.json"
	_, errNotExist := os.Stat(storage_path)
	if os.IsNotExist(errNotExist) {
		file, err := os.Create(storage_path)
		if err != nil {
			jsonrespBody, _ := json.Marshal(structs.HttpResp{
				Status: strconv.Itoa(http.StatusInternalServerError) + ":" + http.StatusText(http.StatusInternalServerError),
				Data:   nil,
			})
			response.Write(jsonrespBody)
			defer file.Close()
			return
		}

		jsonrespBody, _ := json.Marshal(structs.HttpResp{
			Status: strconv.Itoa(http.StatusOK) + ":" + http.StatusText(http.StatusOK),
			Data:   nil,
		})
		response.Write(jsonrespBody)
		defer file.Close()
	}
	itemList := []Item{}

	price, _ := strconv.ParseFloat(request.FormValue("Price"), 64)
	item := Item{
		Id:          request.FormValue("Id"),
		Name:        request.FormValue("Name"),
		Category:    request.FormValue("Category"),
		Price:       price,
		Description: request.FormValue("Description"),
		Image:       request.FormValue("Image"),
	}

	byteData, err := os.ReadFile(storage_path)
	if err != nil {
		fmt.Println(err.Error())
		response.WriteHeader(http.StatusNotFound)
		return
	}
	json.Unmarshal(byteData, &itemList)
	for i, currentData := range itemList {
		if item.Id == currentData.Id {
			if item == currentData {
				response.WriteHeader(http.StatusNoContent) //200SUCCESS
				respBody, _ := json.Marshal(structs.HttpResp{
					Status: strconv.Itoa(http.StatusNoContent) + ":" + http.StatusText(http.StatusNoContent),
					Data:   item,
				})
				response.Write(respBody)
				return
			}
			response.WriteHeader(http.StatusOK)
			itemList[i] = item
			data, _ := json.MarshalIndent(itemList, " ", " ")
			file, _ := os.OpenFile(storage_path, os.O_RDWR, 0644)
			file.Write(data)
			respBody, _ := json.Marshal(structs.HttpResp{
				Status: strconv.Itoa(http.StatusOK) + ":" + http.StatusText(http.StatusOK),
				Data:   data,
			})
			response.Write(respBody)
			return
		}
	}
	response.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(response, "Data is not found")
}

func ShowSelectedItem(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var data []Item
	var selectedData []Item

	file_path := "./assets/item_list.json"
	file, errReadFile := os.ReadFile(file_path)

	if errReadFile != nil {
		response.WriteHeader(http.StatusNotFound)
	}
	json.Unmarshal(file, &data)

	query_cat := request.URL.Query().Get("Category")
	query_price, _ := strconv.ParseFloat(request.URL.Query().Get("Price"), 64)
	query_priceGreater, _ := strconv.ParseBool(request.URL.Query().Get("PriceGreater"))

	for i, v := range data {
		if v.Category == query_cat {
			if (query_priceGreater) && (v.Price > query_price) {
				selectedData = append(selectedData, data[i])
			} else if (!query_priceGreater) && (v.Price < query_price) {
				selectedData = append(selectedData, data[i])
			}
		}
	}
	fmt.Println(selectedData)
	fmt.Fprintln(response, selectedData)
}

func ShowItemDetails(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param_value := params.ByName("id")
	storage_path := "./assets/item_list.json"
	_, errNotExist := os.Stat(storage_path)

	if os.IsNotExist(errNotExist) {
		file, err := os.Create(storage_path)
		if err != nil {
			jsonrespBody, _ := json.Marshal(structs.HttpResp{
				Status: strconv.Itoa(http.StatusInternalServerError) + ":" + http.StatusText(http.StatusInternalServerError),
				Data:   nil,
			})
			response.Write(jsonrespBody)
			defer file.Close()
			return
		}

		jsonrespBody, _ := json.Marshal(structs.HttpResp{
			Status: strconv.Itoa(http.StatusOK) + ":" + http.StatusText(http.StatusOK),
			Data:   nil,
		})
		response.Write(jsonrespBody)
		defer file.Close()
	}
	itemList := []Item{}
	byteData, err := os.ReadFile(storage_path)
	if err != nil {
		fmt.Println(err.Error())
		response.WriteHeader(http.StatusNotFound)
		return
	}
	json.Unmarshal(byteData, &itemList)
	for _, value := range itemList {
		if param_value == value.Id {
			respBody, _ := json.Marshal(structs.HttpResp{
				Status: strconv.Itoa(http.StatusOK) + ":" + http.StatusText(http.StatusOK),
				Data:   value,
			})
			response.Write(respBody)
			return
		}
	}
}
