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
	router.POST("/schedules", handler.NewSchedule)
	router.GET("/schedules", handler.ShowSchedule)
	router.PUT("/schedules/:schedule-id", handler.EditSchedule)
	router.DELETE("/schedules/:schedule-id", handler.DeleteSchedule)
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
