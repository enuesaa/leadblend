package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverCreateCometArgs struct {
	Data string
}

func (r *Resolver) CreateComet(args resolverCreateCometArgs) (*string, error) {
	stoneSrv := service.NewStoneService(r.repos)

	params := dbq.CreateStoneParams{
		Data: args.Data,
	}
	id, err := stoneSrv.Create(params)
	if err != nil {
		return nil, err
	}
	// list patterns
	// if pattern matched, append pattern_id to comit.

	return &id, nil
}
