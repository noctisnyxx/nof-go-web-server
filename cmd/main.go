package main

import (
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
	MuxHandler.HandleFunc("/showorder/", handler.MakeOrder)
	MuxHandler.HandleFunc("/", handler.ShowHomePage)

	webServer := http.Server{
		Addr:    ":8080",
		Handler: MuxHandler,
	}
	err := webServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
