package endpoints

import (
	"net/http"
	"../storage"
)

type Out struct {
	Endpoint    string
	Storage		*storage.Storage
}

func (ch Out) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	//queryValues := req.URL.Query()

	//age := utils.StrToInt(queryValues.Get("age"))
	//gender := utils.StrToInt(queryValues.Get("gender"))
	//showroomId := utils.StrToInt(queryValues.Get("showroomId"))

	//ch.Storage.PersonOutShowroom(showroomId, storage.Person{Age:age, Gender:gender})
	writer.WriteHeader(http.StatusOK)
}

func (ch Out) GetEndpoint() string {
	return ch.Endpoint
}
