package handler

import (
	"fmt"
	"net/http"
)

func ShowHomePage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome to the home page!")
}
