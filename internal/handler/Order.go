package handler

import (
	"fmt"
	"net/http"
)

func ShowOrder(write http.ResponseWriter, request *http.Request) {
	fmt.Fprint(write, "this page desired to be used in displaying the order")
}

func MakeOrder(write http.ResponseWriter, request *http.Request) {
}
