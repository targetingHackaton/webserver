package endpoints

import (
	"net/http"
	"../storage"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

type Front struct {
	Endpoint   string
	Storage    *storage.Storage
	DriverPool *bolt.DriverPool
}

func (ch Front) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	message := "Front"
	writer.Write([]byte(message))
}

func (ch Front) GetEndpoint() string {
	return ch.Endpoint
}
