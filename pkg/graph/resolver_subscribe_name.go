package graph

import (
	"context"
	"fmt"
	"time"
)

func (*Resolver) SubscribeName(ctx context.Context) (<-chan *string, error) {
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

	return ch, nil
}
