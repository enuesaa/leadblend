package graph

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve(repos repository.Repos) error {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3001"},
	}))

	app.POST("/graphql", echo.WrapHandler(&relay.Handler{Schema: gqschema}))
	app.GET("/graphql", wshandle)
	app.GET("/graphql/playground", echo.WrapHandler(playground.Handler("graphql", "/graphql")))

	return app.Start(":3000")
}
