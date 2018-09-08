package endpoints

import (
	"net/http"
	"../storage"
	"utils"
)

type In struct {
	Endpoint    string
	Storage	    *storage.Storage
}

func (ch In) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()

	age := utils.StrToInt(queryValues.Get("age"))
	gender := queryValues.Get("gender")
	showroomId := utils.StrToInt(queryValues.Get("showroomId"))

	ch.Storage.PersonInShowroom(showroomId, storage.Person{AgeIdentifier:age, Gender:gender})
	writer.WriteHeader(http.StatusOK)
}

func (ch In) GetEndpoint() string {
	return ch.Endpoint
}
