package graph

import (
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverDeleteStoneArgs struct {
	Id string
}

func (r *Resolver) DeleteStone(args resolverDeleteStoneArgs) (*bool, error) {
	stoneSrv := service.NewStoneService(r.repos)

	if err := stoneSrv.Delete(args.Id); err != nil {
		return nil, err
	}
	result := true
	return &result, nil
}
