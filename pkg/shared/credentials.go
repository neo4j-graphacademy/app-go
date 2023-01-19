package shared

import "os"

type Neo4jCredentials struct {
	Uri      string
	Username string
	Password string
}

func GetNeo4jCredentials() Neo4jCredentials {
	return Neo4jCredentials{
		Uri:      os.Getenv("NEO4J_URI"),
		Username: os.Getenv("NEO4J_USERNAME"),
		Password: os.Getenv("NEO4J_PASSWORD"),
	}
}
