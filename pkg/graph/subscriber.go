package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/graph-gophers/graphql-go"
)

type SubscribeRequestMsg struct {
	Type    string                     `json:"type"`
	Id      string                     `json:"id"`
	Payload SubscribeRequestMsgPayload `json:"payload"`
}
type SubscribeRequestMsgPayload struct {
	Query string `json:"query"`
}
type SubscribeResponseMsg struct {
	Type    string            `json:"type"`
	Id      string            `json:"id,omitempty"`
	Payload *graphql.Response `json:"payload,omitempty"`
}

type Subscriber struct {
	ws       *websocket.Conn
	gqschema *graphql.Schema
}

func (s *Subscriber) Init(w http.ResponseWriter, r *http.Request, gqschema *graphql.Schema) error {
	headers := http.Header{}
	headers.Add("Sec-Websocket-Protocol", "graphql-transport-ws")

	var upgrader = websocket.Upgrader{}
	ws, err := upgrader.Upgrade(w, r, headers)
	if err != nil {
		return err
	}
	s.ws = ws
	s.gqschema = gqschema

	return nil
}

func (s *Subscriber) checkWsConn() error {
	if s.ws == nil {
		return fmt.Errorf("ws connection does not exist.")
	}
	return nil
}

func (s *Subscriber) Close() error {
	if s.ws == nil {
		return nil
	}
	err := s.ws.Close()
	s.ws = nil
	return err
}

func (s *Subscriber) Read() (SubscribeRequestMsg, error) {
	var req SubscribeRequestMsg
	if err := s.checkWsConn(); err != nil {
		return req, err
	}
	_, reqbyte, err := s.ws.ReadMessage()
	if err != nil {
		return req, err
	}
	if err := json.Unmarshal(reqbyte, &req); err != nil {
		return req, err
	}
	return req, err
}

func (s *Subscriber) Send(msg interface{}) error {
	if err := s.checkWsConn(); err != nil {
		return err
	}
	return s.ws.WriteJSON(msg)
}

func (s *Subscriber) sendAwk() error {
	return s.Send(SubscribeResponseMsg{
		Type: "connection_ack",
	})
}

func (s *Subscriber) subscribe(req SubscribeRequestMsg) error {
	ctx := context.Background()
	ch, err := s.gqschema.Subscribe(ctx, req.Payload.Query, "", nil)
	if err != nil {
		return err
	}
	for {
		data, _ := <-ch
		res := SubscribeResponseMsg{
			Type:    "next",
			Payload: data.(*graphql.Response),
			Id:      req.Id,
		}
		if err := s.checkWsConn(); err != nil {
			break
		}
		if err := s.ws.WriteJSON(res); err != nil {
			return err
		}
	}
	return nil
}

func (s *Subscriber) WaitMessage() error {
	for {
		req, err := s.Read()
		if err != nil {
			return err
		}
		switch req.Type {
		case "connection_init":
			if err := s.sendAwk(); err != nil {
				return err
			}
		case "subscribe":
			if err := s.subscribe(req); err != nil {
				return err
			}
		case "complete":
			if err := s.Close(); err != nil {
				return err
			}
		}
	}
}
