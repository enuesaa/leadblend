package graph

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/graph-gophers/graphql-go"
	"github.com/labstack/echo/v4"
)

type SubscribeRequest struct {
	Type    string `json:"type"`
	Id      string `json:"id"`
	Payload SubscribeRequestPayload `json:"payload"`
}
type SubscribeRequestPayload struct {
	Query string `json:"query"`
}
type SubscribeResponse struct {
	Type    string `json:"type"`
	Id      string `json:"id,omitempty"`
	Payload *graphql.Response `json:"payload,omitempty"`
}

// see https://fits.hatenablog.com/entry/2021/01/12/210926
// see https://docs.aws.amazon.com/ja_jp/appsync/latest/devguide/real-time-websocket-client.html
func handleSubscribe(c echo.Context) error {
	headers := http.Header{}
	headers.Add("Sec-Websocket-Protocol", "graphql-transport-ws")

	var upgrader = websocket.Upgrader{}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), headers)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		_, reqbyte, err := ws.ReadMessage()
		if err != nil {
			return err
		}
		var req SubscribeRequest
		if err := json.Unmarshal(reqbyte, &req); err != nil {
			return err
		}
		if req.Type == "connection_init" {
			res := SubscribeResponse {
				Type: "connection_ack",
			}
			if err := ws.WriteJSON(res); err != nil {
				return err
			}
		} else {
			ctx := context.Background()
			ch, err := gqschema.Subscribe(ctx, req.Payload.Query, "", nil)
			if err != nil {
				return err
			}
			for {
				data, _ := <-ch
				res := SubscribeResponse {
					Type: "next",
					Payload: data.(*graphql.Response),
					Id: req.Id,
				}
				if err := ws.WriteJSON(res); err != nil {
					return err
				}
			}
		}
	}
}
