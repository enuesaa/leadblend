package graph

import (
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverDeleteCometArgs struct {
	Id string
}

func (r *Resolver) DeleteComet(args resolverDeleteCometArgs) (*bool, error) {
	stoneSrv := service.NewStoneService(r.repos)

	if err := stoneSrv.Delete(args.Id); err != nil {
		return nil, err
	}
	result := true
	return &result, nil
}
