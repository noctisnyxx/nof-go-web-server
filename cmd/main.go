package main

import (
	"fmt"
	"net/http"
	"nof-go-web-server/internal/handler"
)

func main() {
	//handler function hanya dapat satu endpoint saja, untuk multiple dapat menggunakan ServeMux
	// var handler1 http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	// 	fmt.Fprintf(writer, "Hello! this is the first handler!\n")
	// 	fmt.Fprintf(writer, "<- writer\n")
	// 	fmt.Fprintf(writer, request.Host+"\n")
	// 	fmt.Fprintf(writer, request.Method+"\n")

	// }

	MuxHandler := http.NewServeMux()
	MuxHandler.HandleFunc("/showorder/", handler.ShowOrder)
	MuxHandler.HandleFunc("/makeorder/", handler.MakeOrder)
	MuxHandler.HandleFunc("/", handler.ShowHomePage)
	MuxHandler.HandleFunc("/reqhandlertest/", handler.ShowRequestTestHandler)

	webServer := http.Server{
		Addr:    ":8080",
		Handler: MuxHandler,
	}
	fmt.Println("the server might be running at " + "http://localhost" + ":8080")
	err := webServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
