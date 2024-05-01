package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/enuesaa/leadblend/pkg/service"
)

type CreatePlanetArgs struct {
	Name    string
	Comment string
}

func (r *Resolver) CreatePlanet(args CreatePlanetArgs) (*string, error) {
	planetSrv := service.NewPlanetService(r.repos)

	params := dbq.CreatePlanetParams{
		Name:    args.Name,
		Comment: args.Comment,
	}
	id, err := planetSrv.Create(params)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
