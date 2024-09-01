package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type order struct {
	Id        string    `json:"id"`
	ItemName  string    `json:"item_name"`
	Price     float64   `json:"item_price"`
	CreatedAt time.Time `json:"order_created_at"`
}

func ShowOrder(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(response, "error: requested method is not allowed")
		return
	}
	fmt.Fprintln(response, "this page desired to be used in displaying the order")
}

func MakeANewOrder(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(strconv.Itoa(http.StatusMethodNotAllowed) + ": requested method is not allowed"))
		return
	}
	var orderList []order
	var orderListJSON string
	newOrder := order{
		Id:        "sampleId",
		ItemName:  "sampleItemName",
		Price:     0.00,
		CreatedAt: time.Now(),
	}

	orderList = append(orderList, newOrder)
	orderListMarshaled, _ := json.Marshal(orderList)
	orderListJSON = string(orderListMarshaled)
	response.Write(orderListMarshaled)

	fmt.Fprintln(response, newOrder)
	fmt.Println(newOrder)
	fmt.Println(orderListJSON)
}
