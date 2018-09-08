package endpoints

import (
	"net/http"
	"../storage"
)

type Camera struct {
	Endpoint   string
	Storage	   *storage.Storage
}

func (ch Camera) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	//queryValues := req.URL.Query()

	//age := utils.StrToInt(queryValues.Get("age"))
	//gender := utils.StrToInt(queryValues.Get("gender"))
	//cameraId := utils.StrToInt(queryValues.Get("cameraId"))
	//showroomId := utils.StrToInt(queryValues.Get("showroomId"))

	//ch.Storage.PersonInFrontOfCamera(showroomId, cameraId, storage.Person{Age:age, Gender:gender})
	writer.WriteHeader(http.StatusOK)
}

func (ch Camera) GetEndpoint() string {
	return ch.Endpoint
}

