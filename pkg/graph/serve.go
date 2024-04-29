package graph

import (
	_ "embed"

	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/99designs/gqlgen/graphql/playground"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	_ "github.com/graph-gophers/graphql-transport-ws/graphqlws"
)

//go:embed query.gql
var schema string

func Serve(repos repository.Repos) error {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3001"},
	}))

	gqschema := graphql.MustParseSchema(schema, &Resolver{})
	app.POST("/graphql", echo.WrapHandler(&relay.Handler{Schema: gqschema}))
	// app.Any("/graphql", echo.WrapHandler(
	// 	graphqlws.NewHandler(gqschema, &relay.Handler{Schema: gqschema}),
	// ))
	app.GET("/graphql/playground", echo.WrapHandler(playground.Handler("graphql", "/graphql")))

	return app.Start(":3000")
}
