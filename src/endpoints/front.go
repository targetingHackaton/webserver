package endpoints

import (
	"net/http"
	"../storage"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"utils"
)

type Front struct {
	Endpoint   string
	Storage    *storage.Storage
	DriverPool *bolt.DriverPool
}

func (ch Front) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()

	age := utils.StrToInt(queryValues.Get("age"))
	gender := queryValues.Get("gender")
	cameraId := utils.StrToInt(queryValues.Get("cameraId"))
	showroomId := utils.StrToInt(queryValues.Get("showroomId"))

	person := storage.Person{AgeIdentifier:age, Gender:gender}
	ch.Storage.PersonInFrontOfCamera(showroomId, cameraId, person)

	writer.WriteHeader(http.StatusOK)
}

func (ch Front) GetEndpoint() string {
	return ch.Endpoint
}
