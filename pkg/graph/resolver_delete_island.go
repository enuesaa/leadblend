package graph

import (
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverDeleteIslandArgs struct {
	Id string
}

func (r *Resolver) DeleteIsland(args resolverDeleteIslandArgs) (*bool, error) {
	islandSrv := service.NewIslandService(r.repos)

	if err := islandSrv.Delete(args.Id); err != nil {
		return nil, err
	}
	result := true
	return &result, nil
}
