package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type SubscribeMessage struct {
	Id      string `json:"id"`
	Payload struct {
		Query string `json:"query"`
	} `json:"payload"`
}

var upgrader = websocket.Upgrader{}

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
