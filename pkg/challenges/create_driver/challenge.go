package main

// Import the driver ("github.com/neo4j/neo4j-go-driver/v5/neo4j")
import (
	"context"
	. "github.com/neo4j-graphacademy/neoflix/pkg/shared"
)

func main() {
	credentials := GetNeo4jCredentials()
	ctx := context.Background()

	ctx.Value(credentials) // TODO: remove this when you start

	// TODO: Create a DriverWithContext Instance
	// TODO: close the driver with defer PanicOnClosureError(ctx, driver)
	// TODO: Open a new Session
	// TODO: close the session with defer PanicOnClosureError(ctx, driver)
	// TODO: Run a Cypher statement
	// TODO: Log the Director value of the first record
}
