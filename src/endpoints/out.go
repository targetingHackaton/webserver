package endpoints

import (
	"../storage"
	"../utils"
	"net/http"
)

type Out struct {
	Endpoint string
	Storage  *storage.Storage
}

func (ch Out) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()

	age := utils.StrToInt(queryValues.Get("age"))
	gender := queryValues.Get("gender")
	showroomId := utils.StrToInt(queryValues.Get("showroomId"))

	ch.Storage.PersonOutShowroom(showroomId, storage.Person{AgeIdentifier: age, Gender: gender})
	writer.WriteHeader(http.StatusOK)
}

func (ch Out) GetEndpoint() string {
	return ch.Endpoint
}
