package neo4j

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func GetNeo4JDriverPool(dsn string, maxConnection int) bolt.DriverPool {
	driverPool, err := bolt.NewDriverPool(dsn, maxConnection)
	if (err != nil) {
		panic(err)
	}

	return driverPool
}
