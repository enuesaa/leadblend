package graph

import "github.com/enuesaa/leadblend/pkg/service"

type resolverIslandArgs struct {
	Id *string
}
func (r *Resolver) Island(args resolverIslandArgs) (*Island, error) {
	islandSrv := service.NewIslandService(r.repos)
	island, err := islandSrv.Get(*args.Id)
	if err != nil {
		return nil, err
	}
	return &Island{island}, nil
}
