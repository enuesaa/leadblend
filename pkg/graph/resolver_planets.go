package graph

import (
	"context"
)

func (r *Resolver) Planets() ([]*Planet, error) {
	list := make([]*Planet, 0)

	if err := r.repos.DB.Open(r.db); err != nil {
		return list, err
	}

	query, err := r.repos.DB.Query()
	if err != nil {
		return list, err
	}
	planets, err := query.ListPlanets(context.Background())
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
