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

/*

// tag::executeread[]
personNode, err := neo4j.ExecuteRead[neo4j.Node](
	ctx,  // <1>
	session, // <2>
	func(tx neo4j.ManagedTransaction) (neo4j.Node, error) { // <3>
		result, err := tx.Run(ctx, cypher, params)
		if err != nil {
			return *new(neo4j.Node), err
		}
		// extract the single result
		// and return it as a neo4j.Node
		return neo4j.SingleTWithContext(ctx, result,  // <4>
			func(record *neo4j.Record) (neo4j.Node, error) {
				node, _, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
				return node, err
			}
		)
	}
)
// end::executeread[]

// tag::executewrite[]
personNode, err := neo4j.ExecuteRead[neo4j.Node](
	ctx,  // <1>
	session, // <2>
	func(tx neo4j.ManagedTransaction) (neo4j.Node, error) { // <3>
		result, err := tx.Run(ctx, cypher, params)
		if err != nil {
			return *new(neo4j.Node), err
		}
		// extract the single result
		// and return it as a neo4j.Node
		return neo4j.SingleTWithContext(ctx, result,  // <4>
			func(record *neo4j.Record) (neo4j.Node, error) {
				node, _, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
				return node, err
			}
		)
	}
)
// end::executewrite[]

# Shortform examples:

// tag::Single[]
// Get the first and only result from the stream.
first, err := record.Single()
// end::Single[]

// tag::Next[]
// .Next() returns false upon error
for result.Next() {
    record := result.Record()
    handleRecord(record)
}
// Err returns the error that caused Next to return false
if err = result.Err(); err != nil {
    handleError(err)
}
// end::Next[]

// tag::NextRecord[]
for result.NextRecord(&record) {
    fmf.Println(record.Keys)
}
// end::NextRecord[]

// tag::Consume[]
summary := result.Consume()
// Time in milliseconds before receiving the first result
fmt.Println(summary.ResultAvailableAfter())
// Time in milliseconds once the final result was consumed
fmt.Println(summary.ResultConsumedAfter())
// end::Consume[]

// tag::Collect[]
remaining, remainingErr := result.Collect()
// end::Collect[]

// tag::keys[]
fmt.Println(record.Keys) // ['p', 'r', 'm']
// end::keys[]

// tag::index[]
// Access a value by its index
fmt.Println(record.Values[0].(neo4j.Node)) // The Person node
// end::index[]

// tag::alias[]
// Access a value by its alias
movie, _ := record.Get("movie")
movieNode := movie.(neo4j.Node)
fmt.Println(movieNode) // The Movie node
// end::alias[]

*/
