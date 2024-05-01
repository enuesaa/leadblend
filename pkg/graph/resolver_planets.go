package graph

import (
	"github.com/enuesaa/leadblend/pkg/service"
)

func (r *Resolver) Planets() ([]*Planet, error) {
	list := make([]*Planet, 0)

	planetSrv := service.NewPlanetService(r.repos)
	planets, err := planetSrv.List()
	if err != nil {
		return list, err
	}

	for _, planet := range planets {
		list = append(list, &Planet{
			id: planet.ID,
			name: planet.Name,
		})
	}

	return list, nil
}
