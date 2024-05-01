package graph

import (
	"context"
	// "fmt"
	// "time"
)

var subscribeEventCh = make(chan *Event)

func (r *Resolver) SubscribeEvent(ctx context.Context) (<-chan *Event, error) {
	// ch := make(chan *Event)

	// go func() {
	// 	defer close(ch)
	// 	for {
	// 		time.Sleep(1 * time.Second)
	// 		name := "a"
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Println("Subscription Closed")
	// 			return
	// 		case ch <- &name:
	// 		}
	// 	}
	// }()

	return subscribeEventCh, nil
}
