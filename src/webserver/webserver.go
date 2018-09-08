package main

import (
	"net/http"
	"../endpoints"
	"../storage"
	"../neo4j"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"os"
	"fmt"
)

const neo4jMaxConnections = 50

type HTTPHandler interface {
	http.Handler
	GetEndpoint() string
}

func main() {
	var dataStorage = storage.Storage{}

	if len(os.Args) != 2 {
		panic("Error")
	}

	var neo4jDsn = os.Args[1]
	fmt.Println("Neo4jDsn: " + neo4jDsn)

	forever := make(chan interface{}, 1)
	neo4jDriverPool := neo4j.GetNeo4JDriverPool(neo4jDsn, neo4jMaxConnections)

	srv := &http.Server{Addr: ":8080"}
	for _, handler := range getHandlers(&dataStorage, &neo4jDriverPool) {
		http.Handle(handler.GetEndpoint(), handler)
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	<-forever
}

func getHandlers(dataStorage *storage.Storage, neo4jDriverPool *bolt.DriverPool) []HTTPHandler {
	return []HTTPHandler{
		endpoints.All{Endpoint: "/all", Storage: dataStorage, DriverPool: neo4jDriverPool},
		endpoints.Person{Endpoint: "/person", Storage: dataStorage, DriverPool: neo4jDriverPool},
		endpoints.Front{Endpoint: "/front", Storage: dataStorage, DriverPool: neo4jDriverPool},
		endpoints.In{Endpoint: "/in", Storage: dataStorage},
		endpoints.Out{Endpoint: "/out", Storage: dataStorage},
		endpoints.Camera{Endpoint: "/camera", Storage: dataStorage},
	}
}
