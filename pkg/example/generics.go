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

// tag::struct[]
type personActedInMovie struct {
	person  neo4j.Node
	actedIn neo4j.Relationship
	movie   neo4j.Node
}

// end::struct[]

/*
// tag::collecttwithcontext[]
neo4j.CollectTWithContext(
	ctx, // <1>
	result, // <2>
	func(record *neo4j.Record) (T any, error) { // <3>
		// Use `neo4j.GetRecordValue` to access values
	},
)
// end::collecttwithcontext[]
*/

// tag::example[]
func collectTExample() {
	// Create a Driver Instance
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		"neo4j+s://dbhash.databases.neo4j.io",    // <1>
		neo4j.BasicAuth("neo4j", "letmein!", ""), // <2>
	)
	PanicOnErr(err)
	defer PanicOnClosureError(ctx, driver)

	// Open a new Session
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer PanicOnClosureError(ctx, session)

	// Define Query
	cypher := `
		MATCH (person:Person)-[actedIn:ACTED_IN]->(movie:Movie {title: $title})
		RETURN person, actedIn, movie
	`
	params := map[string]any{"title": "The Matrix"}

	// tag::use-struct[]
	people, err := neo4j.ExecuteRead(
		ctx,
		session,
		func(tx neo4j.ManagedTransaction) ([]personActedInMovie, error) {
			result, err := tx.Run(ctx, cypher, params)
			if err != nil {
				return nil, err
			}
			return neo4j.CollectTWithContext(
				ctx,
				result,
				func(record *neo4j.Record) (personActedInMovie, error) {
					// tag::getrecordvalue[]
					person, isNil, err := neo4j.GetRecordValue[neo4j.Node](record, "person")
					// end::getrecordvalue[]

					if isNil {
						fmt.Println("person value is nil")
					}

					if err != nil {
						return personActedInMovie{}, err
					}

					actedIn, _, err := neo4j.GetRecordValue[neo4j.Relationship](record, "actedIn")
					if err != nil {
						return personActedInMovie{}, err
					}

					movie, _, err := neo4j.GetRecordValue[neo4j.Node](record, "movie")
					if err != nil {
						return personActedInMovie{}, err
					}

					return personActedInMovie{person, actedIn, movie}, nil
				},
			)
		},
	)
	// end::use-struct[]

	fmt.Println(people)

	// First Row
	first := people[0]

	node := first.person
	rel := first.actedIn

	// tag::node-property[]
	// Get a Node Property
	name, err := neo4j.GetProperty[string](node, "name")
	fmt.Println("Actor name is ", name) // Actor name is Tom Hanks
	// end::node-property[]

	// tag::rel-property[]
	// Get a Relationship Property
	roles, err := neo4j.GetProperty[[]any](rel, "roles")
	fmt.Println("They play ", roles) // They Play ["Woody"]
	// end::rel-property[]
}

// end::example[]

/*
// tag::singletwithcontext[]
neo4j.SingleTWithContext(

	ctx, // <1>
	result, // <2>
	func(record *neo4j.Record) (T any, error) { // <3>
		// Use `neo4j.GetRecordValue` to access values
	},

)
// end::singletwithcontext[]
*/
func singleTExample() {
	// Create a Driver Instance
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		"neo4j+s://dbhash.databases.neo4j.io",    // <1>
		neo4j.BasicAuth("neo4j", "letmein!", ""), // <2>
	)
	PanicOnErr(err)
	defer PanicOnClosureError(ctx, driver)

	// Open a new Session
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer PanicOnClosureError(ctx, session)

	// Define Query
	cypher := `
			MATCH (person:Person)-[:DIRECTED]->(:Movie {title: $title})
			RETURN person.name AS name
		`
	params := map[string]any{"title": "Toy Story"}

	directorName, err := neo4j.ExecuteRead(
		ctx,
		session,
		func(tx neo4j.ManagedTransaction) (string, error) {
			result, err := tx.Run(ctx, cypher, params)
			if err != nil {
				return "", err
			}
			return neo4j.SingleTWithContext(ctx, result,
				func(record *neo4j.Record) (string, error) {
					value, _, err := neo4j.GetRecordValue[string](record, "name")
					return value, err
				})
		})

	fmt.Println(directorName)

}
