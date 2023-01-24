package main

// Import the driver
import (
	"context"

	. "github.com/neo4j-graphacademy/neoflix/pkg/shared"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	// Neo4j Credentials
	credentials := GetNeo4jCredentials()
	// Create a Driver Instance
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		credentials.Uri,
		neo4j.BasicAuth(credentials.Username, credentials.Password, ""),
	)
	PanicOnErr(err)
	defer PanicOnClosureError(ctx, driver)

	// Open a new Session
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer PanicOnClosureError(ctx, session)

	// Execute the `cypher` statement in a write transaction
	cypher := `
		MATCH (m:Movie {title: $title})
		CREATE (p:Person {name: $name})
		CREATE (p)-[:ACTED_IN]->(m)
		RETURN p`
	params := map[string]any{
		"title": "Matrix, The",
		"name":  "Your Name",
	}

	// TODO: Execute the `cypher` statement in a write transaction
	// hint: you can use neo4j.ExecuteWrite[T] over session.ExecuteWrite for better type safety
}
