package main

// Import the driver
import (
	"context"
	"fmt"
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

	// tag::solution[]
	// Execute the `cypher` statement in a write transaction
	cypher := `
		MATCH (m:Movie {title: "Matrix, The"})
		CREATE (p:Person {name: $name})
		CREATE (p)-[:ACTED_IN]->(m)
		RETURN p`
	params := map[string]any{"name": "Your Name"}

	// execute a transaction function with neo4j.ExecuteWrite[T]
	// and properly get the result as a neo4j.Node
	personNode, err := neo4j.ExecuteWrite[neo4j.Node](ctx, session,
		func(tx neo4j.ManagedTransaction) (neo4j.Node, error) {
			result, err := tx.Run(ctx, cypher, params)
			if err != nil {
				return *new(neo4j.Node), err
			}
			// same as before: extract the single result
			// and return it as a neo4j.Node
			return neo4j.SingleTWithContext(ctx, result,
				func(record *neo4j.Record) (neo4j.Node, error) {
					node, _, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
					return node, err
				})
		})
	PanicOnErr(err)
	fmt.Println(personNode)
	// end::solution[]
}
