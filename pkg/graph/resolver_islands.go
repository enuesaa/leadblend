package graph

import "github.com/enuesaa/leadblend/pkg/service"

type resolverIslandsArgs struct {
	PlanetId string
}
func (r *Resolver) Islands(args resolverIslandsArgs) ([]*Island, error) {
	list := make([]*Island, 0)

	islandSrv := service.NewIslandService(r.repos)
	islands, err := islandSrv.List()
	if err != nil {
		return list, err
	}

	for _, island := range islands {
		list = append(list, &Island{
			id: island.ID,
			planetId: island.PlanetID,
			title: island.Title,
			content: island.Title,
			comment: island.Comment,
		})
	}

	return list, nil
}
