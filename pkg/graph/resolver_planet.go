package graph

import (
	"github.com/enuesaa/leadblend/pkg/service"
)

type GetPlanetArgs struct {
	Id string
}
func (r *Resolver) Planet(args GetPlanetArgs) (*Planet, error) {
	planetSrv := service.NewPlanetService(r.repos)
	planet, err := planetSrv.Get(args.Id)
	if err != nil {
		return nil, err
	}
	res := Planet{
		id: planet.ID,
		name: planet.Name,
	}
	return &res, nil
}
