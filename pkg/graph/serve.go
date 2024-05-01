package graph

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo/v4"
)

func HandlePostGraphql(repos repository.Repos) echo.HandlerFunc {
	gqschema := graphql.MustParseSchema(schema, &Resolver{
		repos: repos,
	})
	return echo.WrapHandler(&relay.Handler{Schema: gqschema})
}
func HandleGetGraphql(repos repository.Repos) echo.HandlerFunc {
	gqschema := graphql.MustParseSchema(schema, &Resolver{
		repos: repos,
	})
	handler := func(c echo.Context) error {
		subscriber := Subscriber{}
		if err := subscriber.Init(c.Response(), c.Request(), gqschema); err != nil {
			return err
		}
		// To suppress error: `response.WriteHeader on hijacked connection`
		go func() {
			if err := subscriber.WaitMessage(); err != nil {
				c.Logger().Errorf("Error: %s", err.Error())
			}
		}()
		return nil
	}
	return handler
}

func HandleGraphqlPlayground() echo.HandlerFunc {
	return echo.WrapHandler(playground.Handler("graphql", "/graphql"))
}
