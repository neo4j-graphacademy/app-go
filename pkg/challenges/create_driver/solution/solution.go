package main

// tag::import[]
// Import the driver
import (
	"context"
	"fmt"

	. "github.com/neo4j-graphacademy/neoflix/pkg/shared"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// end::import[]

func main() {
	// tag::credentials[]
	// Neo4j Credentials
	credentials := GetNeo4jCredentials()
	// end::credentials[]

	// Cypher Query and Parameters
	cypher := `
      MATCH (p:Person)-[:DIRECTED]->(:Movie {title: $title})
      RETURN p.name AS Director
    `
	params := map[string]any{"title": "Toy Story"}

	// tag::solution[]
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

	// Run a Cypher statement
	result, err := session.Run(ctx, cypher, params)
	PanicOnErr(err)

	// Log the Director value of the first record
	director, err := neo4j.SingleTWithContext[string](ctx, result,
		// Extract the single record and transform it with a function
		func(record *neo4j.Record) (string, error) {
			// Extract the record value by the specified key
			// and map it to the specified generic type constraint
			director, _, err := neo4j.GetRecordValue[string](record, "Director")
			return director, err
		})
	PanicOnErr(err)
	fmt.Println(director)
	// end::solution[]
}
