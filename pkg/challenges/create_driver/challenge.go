package main

// Import the driver ("github.com/neo4j/neo4j-go-driver/v5/neo4j")
import (
	"context"

	. "github.com/neo4j-graphacademy/neoflix/pkg/shared"
)

func main() {
	// Neo4j Credentials
	credentials := GetNeo4jCredentials()
	ctx := context.Background()

	// Cypher Query and Parameters
	cypher := `
      MATCH (p:Person)-[:DIRECTED]->(:Movie {title: $title})
      RETURN p.name AS Director
    `
	params := map[string]any{"title": "Toy Story"}

	// TODO: Create a DriverWithContext Instance

	// TODO: close the driver with defer PanicOnClosureError(ctx, driver)

	// TODO: Open a new Session

	// TODO: close the session with defer PanicOnClosureError(ctx, driver)

	// TODO: Run a Cypher statement

	// TODO: Log the Director value of the first record
}
