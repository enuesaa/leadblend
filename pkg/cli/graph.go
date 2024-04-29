package cli

import (
	"context"
	_ "embed"
	"fmt"
	"time"

	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"

	"github.com/99designs/gqlgen/graphql/playground"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

//go:embed query.gql
var schema string

type Resolver struct{}

func (*Resolver) Planets() ([]Planet, error) {
	list := make([]Planet, 0)
	list = append(list, Planet{})
	return list, nil
}
func (*Resolver) CreatePlanet(args struct{ Name string }) (*string, error) {
	id := "a"
	return &id, nil
}
func (*Resolver) SubscribeName(ctx context.Context) <-chan *string {
	ch := make(chan *string)

	go func() {
		defer close(ch)
		for {
			time.Sleep(1 * time.Second)
			name := "a"
			select {
			case <-ctx.Done():
				fmt.Println("Subscription Closed")
				return
			case ch <- &name:
			}
		}
	}()

	return ch
}

type Planet struct{}

func (p Planet) Id() graphql.ID {
	return graphql.ID("aaa")
}
func (p Planet) Name() string {
	return "aaaaaa"
}

func CreateGraphCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "graph",
		Short: "graph",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := echo.New()
			app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
				Format: "method=${method}, uri=${uri}, status=${status}\n",
			}))
			app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
				AllowOrigins: []string{"http://localhost:3001"},
			}))

			gqschema := graphql.MustParseSchema(schema, &Resolver{})
			app.POST("/graphql", echo.WrapHandler(&relay.Handler{Schema: gqschema}))
			app.GET("/graphql", func(c echo.Context) error {
				upgrader := websocket.Upgrader{}
				ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
				if err != nil {
					return err
				}
				defer ws.Close()

				// ctx, cancel := context.WithCancel(c.Request().Context())
				// defer cancel()

				// resolver := Resolver{}
				// ch := resolver.SubscribeName(ctx)

				// for {
				// 	select {
				// 	case <-ctx.Done():
				// 		return nil
				// 	case name := <-ch:
				// 		if err := ws.WriteJSON(map[string]interface{}{"name": *name}); err != nil {
				// 			return err
				// 		}
				// 	}
				// }

				return nil
			})
			app.GET("/graphql/playground", echo.WrapHandler(playground.Handler("graphql", "/graphql")))

			return app.Start(":3000")
		},
	}

	return cmd
}
