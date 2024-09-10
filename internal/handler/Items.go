package handler

import (
	"net/http"

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
}

func UpdateItem(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
}

func ShowSelectedItem(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
}

func ShowItemDetails(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
}
