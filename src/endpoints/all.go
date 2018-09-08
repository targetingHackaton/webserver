package endpoints

import (
	"net/http"
	"../storage"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"../utils"
)

type All struct {
	Endpoint   string
	Storage    *storage.Storage
	DriverPool *bolt.DriverPool
}

func (ch All) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	var responseData []int64
	//queryValues := req.URL.Query()
	//showroomId := utils.StrToInt(queryValues.Get("showroomId"))

	neo4jConnection, err := (*ch.DriverPool).OpenPool()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(utils.GetErrorResponse())
		return
	}
	defer neo4jConnection.Close()

	//relevantAge, relevantGender := ch.Storage.GetRelevantAgeAndGender(showroomId)

	cypherQuery := `
		MATCH (c:Customer)-[:ORDERED]->(:Product) WITH c LIMIT 1
			MATCH (c)-[:ORDERED]->(rec:Product)
    		RETURN rec.docId`

	cypherParams := map[string]interface{}{}

	data, err := neo4jConnection.QueryNeo(cypherQuery, cypherParams)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(utils.GetErrorResponse())
		return
	}
	rows, _, err := data.All()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(utils.GetErrorResponse())
		return
	}

	for _, row := range rows {
		responseData = append(responseData, (row[0]).(int64))
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(utils.GetSuccessResponse(responseData))
}

func (ch All) GetEndpoint() string {
	return ch.Endpoint
}

