package graph

import "github.com/enuesaa/leadblend/pkg/service"

type resolverStonesArgs struct {
	IslandId string
}
func (r *Resolver) ListStones(args resolverStonesArgs) ([]*Stone, error) {
	list := make([]*Stone, 0)

	stoneSrv := service.NewStoneService(r.repos)
	stones, err := stoneSrv.ListByIslandId(args.IslandId)
	if err != nil {
		return list, err
	}

	for _, stone := range stones {
		list = append(list, &Stone{stone})
	}

	return list, nil
}
