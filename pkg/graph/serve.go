package graph

import (
	"context"
	_ "embed"
	"encoding/json"
	"net/http"

	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/99designs/gqlgen/graphql/playground"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	"fmt"

	"github.com/gorilla/websocket"
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

	app.POST("/graphql", echo.WrapHandler(&relay.Handler{
		Schema: graphql.MustParseSchema(schema, &Resolver{}),
	}))
	// This is for the graphql subscription.
	// however gophers-graphql does not supported enoughly, so currently not using.
	app.GET("/graphql", wshandle)
	app.GET("/graphql/playground", echo.WrapHandler(playground.Handler("graphql", "/graphql")))

	return app.Start(":3002")
}

var (
	upgrader = websocket.Upgrader{}
)

// see https://fits.hatenablog.com/entry/2021/01/12/210926
// see https://docs.aws.amazon.com/ja_jp/appsync/latest/devguide/real-time-websocket-client.html
func wshandle(c echo.Context) error {
	header := http.Header{}
	header.Add("Sec-Websocket-Protocol", "graphql-transport-ws")
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), header)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
		if string(msg) == `{"type":"connection_init","payload":{}}` {
			if err := ws.WriteMessage(websocket.TextMessage, []byte(`{"type":"connection_ack"}`)); err != nil {
				c.Logger().Error(err)
			}
		} else {
			gqschema :=  graphql.MustParseSchema(schema, &Resolver{})
			type SubscribeMessage struct {
				Id string `json:"id"`
				Payload struct {
					Query string `json:"query"`
				} `json:"payload"`
			}
			var message SubscribeMessage
			if err := json.Unmarshal(msg, &message); err != nil {
				c.Logger().Error(err)
			}
			ch, err := gqschema.Subscribe(context.Background(), message.Payload.Query, "", nil)
			if err != nil {
				c.Logger().Error(err)
			}
			for {
				a, _ := <-ch
				b, _ := json.Marshal(a)
				fmt.Println(string(b))
				res := fmt.Sprintf(`{"payload":%s,"id":"%s","type":"next"}`, b, message.Id)
				if err := ws.WriteMessage(websocket.TextMessage, []byte(res)); err != nil {
					c.Logger().Error(err)
				}
			}
		}
	}
}
