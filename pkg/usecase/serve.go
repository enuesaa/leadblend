package usecase

import (
	"github.com/enuesaa/leadblend/pkg/graph"
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve(repos repository.Repos) error {
	if err := repos.DB.Open(); err != nil {
		return err
	}
	defer repos.DB.Close()

	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3001"},
	}))

	app.POST("/graphql", graph.HandlePostGraphql(repos))
	app.GET("/graphql", graph.HandleGetGraphql(repos))
	app.GET("/graphql/playground", graph.HandleGraphqlPlayground())

	return app.Start(":3000")
}
	