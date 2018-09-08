package endpoints

import (
	"net/http"
	"../storage"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

type Person struct {
	Endpoint   string
	Storage    *storage.Storage
	DriverPool *bolt.DriverPool
}

func (ch Person) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	//var responseData []int64
	//var cypherQuery string
	//var cypherParams map[string]interface{}

	//queryValues := req.URL.Query()
	//showroomId := utils.StrToInt(queryValues.Get("showroomId"))
	//email := utils.StrToInt(queryValues.Get("email"))


	writer.WriteHeader(http.StatusOK)

}

func (ch Person) GetEndpoint() string {
	return ch.Endpoint
}
