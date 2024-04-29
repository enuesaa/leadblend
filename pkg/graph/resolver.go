package graph

import (
	"context"
	"fmt"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

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
