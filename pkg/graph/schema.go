package graph

import (
	_ "embed"

	"github.com/graph-gophers/graphql-go"
)

//go:embed schema.gql
var schema string
var gqschema = graphql.MustParseSchema(schema, &Resolver{})
