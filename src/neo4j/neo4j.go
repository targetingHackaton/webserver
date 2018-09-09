package neo4j

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"math/rand"
)

const randomSkipMax = 1000

func GetNeo4JDriverPool(dsn string, maxConnection int) bolt.DriverPool {
	driverPool, err := bolt.NewDriverPool(dsn, maxConnection)
	if (err != nil) {
		panic(err)
	}

	return driverPool
}

func GetFallbackScenario(neo4jConn bolt.Conn) []int64 {
	var responseData []int64

	cypherQuery := `
	MATCH (c:Customer)-[:ORDERED]->(:Product)<-[:IS_MAIN_VENDOR]-(:Vendor{vendorId:1})
		WITH c SKIP {randomSkip} LIMIT 300
			MATCH (c)-[:ORDERED]->(p:Product)<-[:IS_MAIN_VENDOR]-(:Vendor{vendorId:1})
    		WHERE p.available = true AND p.sensible = false
    		WITH p.docId AS DOCID, count(c) AS freq
    		ORDER BY freq DESC
 		LIMIT 20
    	RETURN DOCID
	`
	cypherParams := map[string]interface{}{
		"randomSkip": rand.Intn(randomSkipMax),
	}

	data, err := neo4jConn.QueryNeo(cypherQuery, cypherParams)

	if err != nil {
		return responseData
	}

	rows, _, err := data.All()
	if err != nil {
		return responseData
	}

	for _, row := range rows {
		responseData = append(responseData, (row[0]).(int64))
	}

	return responseData
}
