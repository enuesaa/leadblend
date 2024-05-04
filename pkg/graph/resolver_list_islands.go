package graph

import (
	"fmt"

	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverIslandsArgs struct {
	PlanetId string
}

func (r *Resolver) ListIslands(args resolverIslandsArgs) ([]*Island, error) {
	list := make([]*Island, 0)

	islandSrv := service.NewIslandService(r.repos)
	fmt.Println(args.PlanetId)
	islands, err := islandSrv.ListByPlanetId(args.PlanetId)
	if err != nil {
		return list, err
	}

	for _, island := range islands {
		list = append(list, &Island{island})
	}

	return list, nil
}
