package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

// handling menggunakan library httprouter - julianschmidt

type name struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Account struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	Fullname     name      `json:"fullname"`
	RegisteredAt time.Time `json:"registered_at"`
	UpdatedAt    time.Time `json:"last_update"`
}

func AddAccount(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	accountsProfilePath := "./assets/accounts_profile.json"
	_, err := os.Stat(accountsProfilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(accountsProfilePath)
		if err != nil {
			defer file.Close()
		}
		defer file.Close()
	}
	// data := Account{
	// 	Id:       request.FormValue("Id"),
	// 	Username: request.FormValue("Username"),
	// 	Fullname: name{
	// 		FirstName: request.FormValue("FirstName"),
	// 		LastName:  request.FormValue("LastName"),
	// 	},
	// 	RegisteredAt: time.Now(),
	// 	UpdatedAt:    time.Now(),
	// }
}

func ShowAccount(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}

func DeleteAccount(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}

func UpdateAccoutn(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}
