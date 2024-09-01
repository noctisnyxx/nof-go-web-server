package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Item struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"picture_link"`
}

var itemList []Item

func AddItem(response http.ResponseWriter, request *http.Request) {
	storage_path := "./assets/item_list.json"
	_, errChecking := os.Stat(storage_path)
	if os.IsNotExist(errChecking) {
		file, err := os.Create(storage_path)
		if err != nil {
			fmt.Println(err.Error())
			defer file.Close()
			return
		}
		defer file.Close()

	}
	switch request.Method {
	case http.MethodPost:
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
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal(byteData, &itemList)

		for _, currentData := range itemList {
			if item.Id == currentData.Id {
				response.WriteHeader(http.StatusConflict)
				fmt.Fprintln(response, "same Id is not allowed")
				return
			}
		}

		file, err := os.OpenFile(storage_path, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println(err.Error())
			response.WriteHeader(http.StatusNotModified)
			return
		}

		itemList = append(itemList, item)
		data, _ := json.MarshalIndent(itemList, " ", " ")
		file.Write(data)

	case http.MethodPut:
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
					fmt.Println("There is no update, the previous and the new are the same.")
					response.WriteHeader(http.StatusNoContent)
					return
				}
				itemList[i] = item
				data, _ := json.MarshalIndent(itemList, " ", " ")
				file, _ := os.OpenFile(storage_path, os.O_RDWR, 0644)
				file.Write(data)
				fmt.Fprintln(response, "Succesfully update the data!")
				return
			}
		}
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(response, "Data is not found")
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ShowSelectedItem(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var data []Item
	var selectedData []Item

	file_path := "./assets/item_list.json"
	file, errReadFile := os.ReadFile(file_path)
	if errReadFile != nil {
		response.WriteHeader(http.StatusNotFound)
	}
	json.Unmarshal(file, &data)
	fmt.Println(data)

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
