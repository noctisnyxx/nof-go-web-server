package handler

import (
	"fmt"
	"net/http"
)

func ShowRequestTestHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Uji coba isi request"+"\n")
	fmt.Fprintf(response, "Method: "+request.Method+"\n")
	fmt.Fprintf(response, "RequestURI: "+request.RequestURI+"\n")
	query := request.URL.Query().Get("query_data")
	if query == "" {
		fmt.Fprintf(response, "there is no data in query_data")
	} else {
		fmt.Fprintf(response, "query=%s", query)
	}
}
