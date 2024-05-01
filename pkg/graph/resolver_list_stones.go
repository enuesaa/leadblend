package graph

import "github.com/enuesaa/leadblend/pkg/service"

func (r *Resolver) ListStones() ([]*Stone, error) {
	list := make([]*Stone, 0)

	stoneSrv := service.NewStoneService(r.repos)
	stones, err := stoneSrv.List()
	if err != nil {
		return list, err
	}

	for _, stone := range stones {
		list = append(list, &Stone{stone})
	}

	return list, nil
}
