package graph

import (
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverDeletePlanetArgs struct {
	Name string
}

func (r *Resolver) DeletePlanet(args resolverDeletePlanetArgs) (*bool, error) {
	planetSrv := service.NewPlanetService(r.repos)

	if err := planetSrv.DeleteByName(args.Name); err != nil {
		return nil, err
	}
	result := true
	return &result, nil
}
