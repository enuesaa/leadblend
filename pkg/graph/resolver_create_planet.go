package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverCreatePlanetArgs struct {
	Name    string
	Comment string
}

func (r *Resolver) CreatePlanet(args resolverCreatePlanetArgs) (*string, error) {
	planetSrv := service.NewPlanetService(r.repos)

	params := dbq.CreatePlanetParams{
		Name:    args.Name,
		Comment: args.Comment,
	}
	id, err := planetSrv.Create(params)
	if err != nil {
		return nil, err
	}

	go func() {
		subscribeEventCh <- &Event{
			page: "planet",
			userName: "aaa",
		}
	}()

	return &id, nil
}
