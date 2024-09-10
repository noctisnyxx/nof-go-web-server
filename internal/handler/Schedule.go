package handler

import (
	"net/http"
	"nof-go-web-server/internal/module"
	"nof-go-web-server/internal/module/database"
	"nof-go-web-server/internal/module/envs"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewSchedule(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	res := module.HttpBody{
		Status: http.StatusOK,
		Data:   "Success to add a new schedule!",
	}
	db := new(database.Mongo)
	err := db.Connect(envs.MONGO_ATLAS)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Data = "Unable to connect to the database"
	}
	newSch, err := module.HttpRequestBodyJsonReader(request)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Data = "Failed to read the json request body"
	}
	col := db.Client.Database("DQAHotroom").Collection("Schedules")
	if _, err := col.InsertOne(db.Context, newSch, options.InsertOne()); err != nil {
		res.Status = http.StatusInternalServerError
		res.Data = "Failed to add to the database"
	}

}

func EditSchedule(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}

func AbortSchedule(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}

func SearchSchedule(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}
