package endpoints

import (
	"net/http"
	"../storage"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"../utils"
	"../neo4j"
	"math/rand"
)

const randomSkipLimit = 500

type Camera struct {
	Endpoint   string
	Storage	   *storage.Storage
	DriverPool *bolt.DriverPool
}

func (ch Camera) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	var responseData []int64
	var cypherQuery string
	var cypherParams map[string]interface{}
	var rows [][]interface{}
	var err error

	queryValues := req.URL.Query()
	showroomId := utils.StrToInt(queryValues.Get("showroomId"))
	cameraId := utils.StrToInt(queryValues.Get("cameraId"))

	neo4jConnection, err := (*ch.DriverPool).OpenPool()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(utils.GetErrorResponse())
		return
	}
	defer neo4jConnection.Close()

	person, isValid := ch.Storage.GetPersonInFrontOfCamera(showroomId, cameraId)

	if isValid {
		cypherQuery = `
		MATCH (c:Customer {gender:{gender}})-[:ORDERED]->(:Product)<-[:IS_MAIN_VENDOR]-(:Vendor{vendorId:1})
		WHERE {minAge} <= c.age <= {maxAge}
		WITH c SKIP {randomSkip} LIMIT 500
			MATCH (c)-[:ORDERED]->(p:Product)<-[:IS_MAIN_VENDOR]-(:Vendor{vendorId:1})
    		WHERE p.available = true AND p.sensible = false
    		WITH p.docId AS DOCID, count(c) AS freq
    		ORDER BY freq DESC
 		LIMIT 20
    	RETURN DOCID
	`

		cypherParams = map[string]interface{}{
			"gender": person.Gender,
			"minAge": storage.AgeIntervals[person.AgeIdentifier].AgeMin,
			"maxAge": storage.AgeIntervals[person.AgeIdentifier].AgeMax,
			"randomSkip": rand.Intn(randomSkipLimit),
		}

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
		neo4jConnectionFallback, err := (*ch.DriverPool).OpenPool()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(utils.GetErrorResponse())
			return
		}
		defer neo4jConnection.Close()
		responseData = neo4j.GetFallbackScenario(neo4jConnectionFallback)
	} else {
		for _, row := range rows {
			responseData = append(responseData, (row[0]).(int64))
		}
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(utils.GetSuccessResponse(responseData))
}

func (ch Camera) GetEndpoint() string {
	return ch.Endpoint
}

