package endpoints

import (
	"net/http"
	"../storage"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"utils"
	"../neo4j"
)

type Person struct {
	Endpoint   string
	Storage    *storage.Storage
	DriverPool *bolt.DriverPool
}

func (ch Person) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	var responseData []int64
	var cypherQuery string
	var cypherParams map[string]interface{}
	var rows [][]interface{}
	var err error

	queryValues := req.URL.Query()
	email := queryValues.Get("email")

	neo4jConnection, err := (*ch.DriverPool).OpenPool()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(utils.GetErrorResponse())
		return
	}
	defer neo4jConnection.Close()

	cypherQuery = `
		MATCH (c:Customer {email:{email}})-[:ORDERED|VISITED]->(:Product)<-[:ORDERED]-(o:Customer)       		
	WITH c, o LIMIT 500
		MATCH (o)-[:ORDERED]->(p:Product)<-[:IS_MAIN_VENDOR]-(:Vendor{vendorId:1})
		WHERE p.available = true AND p.sensible = false AND NOT (c)-[:ORDERED]->(p)
		WITH p.docId AS DOCID, count(o) AS freq
		ORDER BY freq DESC
		LIMIT 20
	RETURN DOCID
	`

	cypherParams = map[string]interface{}{
		"email": email,
	}

	data, err := neo4jConnection.QueryNeo(cypherQuery, cypherParams)

	if err == nil {
		rows, _, err = data.All()
		if err != nil {
			rows = [][]interface{}{}
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

func (ch Person) GetEndpoint() string {
	return ch.Endpoint
}
