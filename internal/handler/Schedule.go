package handler

import (
	"net/http"
	"nof-go-web-server/internal/database"
	"nof-go-web-server/internal/utils"
	"nof-go-web-server/internal/utils/envs"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewSchedule(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	res := utils.ResponseBody{
		Status: http.StatusOK,
		Data:   "Success to add a new schedule!",
	}
	var newSch utils.ScheduleData
	newSch, err := utils.HttpRequestBodyReader[utils.ScheduleData](request)
	newSch.CreatedAt = time.Now()
	newSch.UpdatedAt = time.Now()
	if err != nil {
		res.UpdateHttpResponse(writer, http.StatusInternalServerError, "Failed to read the json request body")
		return
	}
	if newSch.Title == "" || newSch.TestMode == "" {
		res.UpdateHttpResponse(writer, http.StatusBadRequest, "There are missing values for the required parameters")
		return
	}

	if newSch.Start.IsZero() || newSch.End.IsZero() {
		res.UpdateHttpResponse(writer, http.StatusBadRequest, "Schedule start and end value has not filled yet")
		return
	}

	if newSch.Start.After(newSch.End) || newSch.Start.Equal(newSch.End) {
		res.UpdateHttpResponse(writer, http.StatusBadRequest, "Start schedule is after the stop schedule or has the same value!")
		return
	}

	db := new(database.Mongo)
	if err := db.Connect(envs.MONGO_ATLAS); err != nil {
		res.UpdateHttpResponse(writer, http.StatusInternalServerError, "Unable to connect to the database")
		return
	}
	defer db.CloseClientDB()
	newSch.ScheduleId = uuid.New().String()
	newSch.Status = "Initiating"
	col := db.Client.Database("DQAHotroom").Collection("Schedules")
	if _, err := col.InsertOne(db.Context, newSch, options.InsertOne()); err != nil {
		res.UpdateHttpResponse(writer, http.StatusInternalServerError, "Failed to add the data to the database")
		return
	}
	res.UpdateHttpResponse(writer, res.Status, res.Data)
}

func EditSchedule(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var storedData utils.ScheduleData
	res := utils.ResponseBody{
		Status: http.StatusOK,
		Data:   "Data has been updated",
	}
	//INITIATE THE DATABASE, AND FIND THE COLLECTION

	db := new(database.Mongo)
	db.Connect(envs.MONGO_ATLAS)
	col := db.Client.Database("DQAHotroom").Collection("Schedules")
	//FIND THE DATA
	//1. What is data you are looking for?

	updtSchd, err := utils.HttpRequestBodyReader[utils.ScheduleData](request)
	if err != nil {
		res.UpdateHttpResponse(writer, http.StatusInternalServerError, "Failed to update the data")
	}
	filter := bson.M{
		"schedule_id": updtSchd.ScheduleId,
	}

	if err := col.FindOne(db.Context, filter).Decode(&storedData); err != nil {
		res.UpdateHttpResponse(writer, http.StatusInternalServerError, "Failed to update the data")
	}

	//IF NOT FOUND SEND RESPONSE
	//REPLACE THE INSERTED REQUESTS DATA TO THE DATA

}

func PauseSchedule(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}

func AbortSchedule(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}

func SearchSchedule(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}

func ShowSchedule(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	res := utils.ResponseBody{
		Status: http.StatusOK,
	}
	db := new(database.Mongo)
	if err := db.Connect(envs.MONGO_ATLAS); err != nil {
		res.UpdateHttpResponse(writer, http.StatusInternalServerError, "failed to connect to the database")
		return
	}
	col := db.Client.Database("DQAHotroom").Collection("Schedules")
	filter := bson.M{}
	sort := bson.M{}

	if sort_title := request.URL.Query().Get("sort_title"); sort_title == "asc" {
		sort["title"] = 1
	} else if sort_title == "desc" {
		sort["title"] = -1
	}

	if sort_status := request.URL.Query().Get("sort_status"); sort_status == "asc" {
		sort["status"] = 1
	} else if sort_status == "desc" {
		sort["status"] = -1
	}

	if sort_testmode := request.URL.Query().Get("sort_testmode"); sort_testmode == "asc" {
		sort["test_mode"] = 1
	} else if sort_testmode == "desc" {
		sort["test_mode"] = -1
	}

	if sort_group := request.URL.Query().Get("sort_group"); sort_group == "asc" {
		sort["group"] = 1
	} else if sort_group == "desc" {
		sort["group"] = -1
	}

	if sort_createdat := request.URL.Query().Get("sort_createdat"); sort_createdat == "asc" {
		sort["created_at"] = 1
	} else if sort_createdat == "desc" {
		sort["created_at"] = -1
	}

	if sort_start := request.URL.Query().Get("sort_start"); sort_start == "asc" {
		sort["start"] = 1
	} else if sort_start == "desc" {
		sort["start"] = -1
	}

	if sort_end := request.URL.Query().Get("sort_end"); sort_end == "asc" {
		sort["end"] = 1
	} else if sort_end == "desc" {
		sort["end"] = -1
	}

	if query_status := request.URL.Query().Get("status"); query_status != "" {
		filter["status"] = query_status
	}
	if query_title := request.URL.Query().Get("title"); query_title != "" {
		filter["title"] = query_title
	}
	if query_testMode := request.URL.Query().Get("test_mode"); query_testMode != "" {
		filter["test_mode"] = query_testMode
	}
	if query_group := request.URL.Query().Get("group"); query_group != "" {
		filter["group"] = query_group
	}
	if query_createdat := request.URL.Query().Get("created_at"); query_createdat != "" {
		filter["created_at"] = query_createdat
	}
	if query_updatedat := request.URL.Query().Get("updated_at"); query_updatedat != "" {
		filter["updated_at"] = query_updatedat
	}
	if query_start := request.URL.Query().Get("start_date"); query_start != "" {
		filter["start"] = query_start
	}
	if query_end := request.URL.Query().Get("end_date"); query_end != "" {
		filter["end"] = query_end
	}

	opts := options.Find().SetSort(sort)
	cursor, err := col.Find(db.Context, filter, opts)
	if err != nil {
		res.UpdateHttpResponse(writer, http.StatusInternalServerError, "failed to make a query")
		return
	}
	defer cursor.Close(db.Context)

	var results []bson.M
	if err := cursor.All(db.Context, &results); err != nil {
		res.UpdateHttpResponse(writer, http.StatusInternalServerError, "failed to decode documents")
		return
	}
	res.Data = results
	res.UpdateHttpResponse(writer, res.Status, res.Data)
}
