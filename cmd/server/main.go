package main

import (
	"fmt"
	"net/http"
	"nof-go-web-server/internal/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//handler function hanya dapat satu endpoint saja, untuk multiple dapat menggunakan ServeMux
	router := httprouter.New()
	router.GET("/", handler.ShowHomePage)
	router.POST("/shopkeeper/additem/", handler.AddItem)
	router.PUT("/shopkeeper/additem/", handler.UpdateItem)
	router.GET("/user/showitem/", handler.ShowSelectedItem)
	router.GET("/user/showitem/details/:id", handler.ShowItemDetails)
	router.POST("/newschedule/", handler.NewSchedule)
	router.GET("/showschedule/", handler.ShowSchedule)
	webServer := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("the server might be running at " + "http://localhost" + webServer.Addr)
	err := webServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
