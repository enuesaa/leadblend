package graph

import (
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverLinkCometArgs struct {
	Id string
	IslandId string
}
func (r *Resolver) LinkComet(args resolverLinkCometArgs) (*bool, error) {
	stoneSrv := service.NewStoneService(r.repos)

	if err := stoneSrv.Link(args.Id, args.IslandId); err != nil {
		return nil, err
	}
	res := true
	return &res, nil
}
