package graph

import "github.com/enuesaa/leadblend/pkg/service"

type resolverCometArgs struct {
	Id *string
}

func (r *Resolver) GetComet(args resolverCometArgs) (*Comet, error) {
	stoneSrv := service.NewStoneService(r.repos)
	stone, err := stoneSrv.Get(*args.Id)
	if err != nil {
		return nil, err
	}

	return &Comet{stone}, nil
}
