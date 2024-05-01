package graph

import "github.com/enuesaa/leadblend/pkg/service"


type resolverStoneArgs struct {
	Id *string
}

func (r *Resolver) Stone(args resolverStoneArgs) (*Stone, error) {
	stoneSrv := service.NewStoneService(r.repos)
	stone, err := stoneSrv.Get(*args.Id)
	if err != nil {
		return nil, err
	}

	return &Stone{stone}, nil
}
