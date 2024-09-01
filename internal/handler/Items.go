package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Item struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"picture_link"`
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
	case "POST":
		item := Item{
			request.FormValue("Id"),
			request.FormValue("Name"),
			request.FormValue("Description"),
			request.FormValue("Image"),
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

	case "PUT":
		item := Item{
			request.FormValue("Id"),
			request.FormValue("Name"),
			request.FormValue("Description"),
			request.FormValue("Image"),
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
					response.WriteHeader(http.StatusNotFound)
					return
				}
				itemList[i] = item
				data, _ := json.MarshalIndent(itemList, " ", " ")
				file, _ := os.OpenFile(storage_path, os.O_RDWR, 0644)
				file.Write(data)
				return
			}
		}
		response.WriteHeader(http.StatusNotFound)

		fmt.Fprintln(response, "Data is not found")
		fmt.Println("Sampe sini ga ya")
	}
}
