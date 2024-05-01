package graph

import (
	"fmt"

	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/enuesaa/leadblend/pkg/service"
)

type QueryPlanetArgs struct {
	Id   *string
	Name *string
}

func (r *Resolver) Planet(args QueryPlanetArgs) (*Planet, error) {
	planetSrv := service.NewPlanetService(r.repos)

	var err error
	var planet dbq.Planet
	if args.Id != nil {
		planet, err = planetSrv.Get(*args.Id)
	} else if args.Name != nil {
		planet, err = planetSrv.GetByName(*args.Name)
	} else {
		err = fmt.Errorf("not found.")
	}
	if err != nil {
		return nil, err
	}

	res := Planet{
		id:   planet.ID,
		name: planet.Name,
	}
	return &res, nil
}
