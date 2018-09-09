package endpoints

import (
	"../storage"
	"../utils"
	"net/http"
	"fmt"
	"strconv"
)

type In struct {
	Endpoint string
	Storage  *storage.Storage
}

func (ch In) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()

	age := utils.StrToInt(queryValues.Get("age"))
	gender := queryValues.Get("gender")
	showroomId := utils.StrToInt(queryValues.Get("showroomId"))

	fmt.Println("A intrat cineva de "+strconv.Itoa(age)+ " ani")
	ch.Storage.PersonInShowroom(showroomId, storage.Person{AgeIdentifier: age, Gender: gender})
	writer.WriteHeader(http.StatusOK)
}

func (ch In) GetEndpoint() string {
	return ch.Endpoint
}
