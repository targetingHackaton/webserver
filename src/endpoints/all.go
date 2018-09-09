package endpoints

import (
	"net/http"
	"../storage"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"../utils"
	"fmt"
	"../neo4j"
)

type All struct {
	Endpoint   string
	Storage    *storage.Storage
	DriverPool *bolt.DriverPool
}

func (ch All) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	var responseData []int64
	var cypherQuery string
	var cypherParams map[string]interface{}
	var rows [][]interface{}
	var err error

	queryValues := req.URL.Query()
	showroomId := utils.StrToInt(queryValues.Get("showroomId"))

	neo4jConnection, err := (*ch.DriverPool).OpenPool()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(utils.GetErrorResponse())
		return
	}
	defer neo4jConnection.Close()

	//if storage.Showroom[showroomId]

	relevantAge, relevantGender := ch.Storage.GetRelevantAgeAndGender(showroomId)

	if relevantAge != -1 && relevantGender != "" {
		ageInterval := storage.AgeIntervals[relevantAge]

		cypherQuery = `
		MATCH (c:Customer {gender:{gender}})-[:ORDERED]->(:Product)<-[:IS_MAIN_VENDOR]-(:Vendor{vendorId:1})
		WHERE {minAge} <= c.age <= {maxAge}
		WITH c LIMIT 200
			MATCH (c)-[:ORDERED]->(p:Product)<-[:IS_MAIN_VENDOR]-(:Vendor{vendorId:1})
    		WHERE p.available = true AND p.sensible = false
    		WITH p.docId AS DOCID, count(c) AS freq
    		ORDER BY freq DESC
 		LIMIT 20
    	RETURN DOCID
	`

		cypherParams = map[string]interface{}{
			"gender": relevantGender,
			"minAge": ageInterval.AgeMin,
			"maxAge": ageInterval.AgeMax,
		}

		fmt.Println(cypherParams)

		data, err := neo4jConnection.QueryNeo(cypherQuery, cypherParams)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(utils.GetErrorResponse())
			return
		}
		rows, _, err = data.All()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(utils.GetErrorResponse())
			return
		}
	}

	if len(rows) == 0 {
		responseData = neo4j.GetFallbackScenario(neo4jConnection)
	} else {
		for _, row := range rows {
			responseData = append(responseData, (row[0]).(int64))
		}
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(utils.GetSuccessResponse(responseData))
}

func (ch All) GetEndpoint() string {
	return ch.Endpoint
}

