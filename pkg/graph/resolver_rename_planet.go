package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverRenamePlanetArgs struct {
	Name string
	NewName string
}

func (r *Resolver) RenamePlanet(args resolverRenamePlanetArgs) (*bool, error) {
	planetSrv := service.NewPlanetService(r.repos)

	params := dbq.UpdatePlanetNameParams{
		Name: args.Name,
		NewName: args.NewName,
	}
	if err := planetSrv.Rename(params); err != nil {
		return nil, err
	}
	result := true
	return &result, nil
}
