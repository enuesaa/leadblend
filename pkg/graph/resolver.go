package graph

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Resolver struct {
	repos repository.Repos
}

func Serve(repos repository.Repos) error {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3001"},
	}))

	gqschema := graphql.MustParseSchema(schema, &Resolver{
		repos: repos,
	})
	app.POST("/graphql", echo.WrapHandler(&relay.Handler{Schema: gqschema}))
	app.GET("/graphql", func(c echo.Context) error {
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
	})
	app.GET("/graphql/playground", echo.WrapHandler(playground.Handler("graphql", "/graphql")))

	return app.Start(":3000")
}
