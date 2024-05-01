package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverCreateStoneArgs struct {
	Data string
}

func (r *Resolver) CreateStone(args resolverCreateStoneArgs) (*string, error) {
	stoneSrv := service.NewStoneService(r.repos)

	params := dbq.CreateStoneParams{
		Data: args.Data,
	}
	id, err := stoneSrv.Create(params)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
