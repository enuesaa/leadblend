package graph

import "github.com/enuesaa/leadblend/pkg/service"

func (r *Resolver) ListComets() ([]*Comet, error) {
	list := make([]*Comet, 0)

	stoneSrv := service.NewStoneService(r.repos)
	stones, err := stoneSrv.ListComets()
	if err != nil {
		return list, err
	}

	for _, stone := range stones {
		list = append(list, &Comet{stone})
	}

	return list, nil
}
