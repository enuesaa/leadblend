package graph

import (
	"fmt"

	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverPlanetArgs struct {
	Id   *string
	Name *string
}

func (r *Resolver) GetPlanet(args resolverPlanetArgs) (*Planet, error) {
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
	return &Planet{planet}, nil
}
