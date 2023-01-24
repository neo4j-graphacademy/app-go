package main

// tag::import[]
// Import the driver
import (
	"context"
	"fmt"
	"time"

	. "github.com/neo4j-graphacademy/neoflix/pkg/shared"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// end::import[]

func main() {
	// tag::driver[]
	// Create a Driver Instance
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		"neo4j+s://dbhash.databases.neo4j.io",    // <1>
		neo4j.BasicAuth("neo4j", "letmein!", ""), // <2>
	)
	// end::driver[]
	PanicOnErr(err)
	// tag::close[]
	defer PanicOnClosureError(ctx, driver)
	// end::close[]

	// tag::verify[]
	// Verify Connectivity
	PanicOnErr(driver.VerifyConnectivity(ctx))
	// end::verify[]

	// Open a new Session
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer PanicOnClosureError(ctx, session)

	// tag::oneoff[]
	// Execute a Cypher statement in an auto-commit transaction
	result, err := session.Run(
		ctx, // <1>
		`
		MATCH (p:Person)-[:DIRECTED]->(:Movie {title: $title})
		RETURN p
		`, // <2>
		map[string]any{"title": "The Matrix"}, // <3>
		func(txConfig *neo4j.TransactionConfig) {
			txConfig.Timeout = 3 * time.Second // <4>
		},
	)
	PanicOnErr(err)
	// end::oneoff[]

	// tag::oneoffresult[]
	// Get all Person nodes
	people, err := neo4j.CollectTWithContext[neo4j.Node](ctx, result,
		func(record *neo4j.Record) (neo4j.Node, error) {
			person, _, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
			return person, err
		})
	PanicOnErr(err)
	fmt.Println(people)
	// end::oneoffresult[]

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
}
