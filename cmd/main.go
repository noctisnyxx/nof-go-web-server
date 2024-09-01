package main

import (
	"fmt"
	"net/http"
	"nof-go-web-server/internal/handler"
)

func main() {
	//handler function hanya dapat satu endpoint saja, untuk multiple dapat menggunakan ServeMux
	MuxHandler := http.NewServeMux()
	MuxHandler.HandleFunc("/showorder/", handler.ShowOrder)
	MuxHandler.HandleFunc("/makeorder/", handler.MakeANewOrder)
	MuxHandler.HandleFunc("/", handler.ShowHomePage)
	MuxHandler.HandleFunc("/reqhandlertest/", handler.ShowRequestTestHandler)
	MuxHandler.HandleFunc("/shopkeeper/additem/", handler.AddItem)
	MuxHandler.HandleFunc("/user/showitem/", handler.ShowSelectedItem)

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
