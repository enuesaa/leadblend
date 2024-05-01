package graph

import (
	"context"

	"github.com/enuesaa/leadblend/pkg/repository/dbq"
)

type CreatePlanetArgs struct {
	Name string
}
func (r *Resolver) CreatePlanet(args CreatePlanetArgs) (*string, error) {
	id := "a"

	if err := r.repos.DB.Open(r.db); err != nil {
		return &id, err
	}

	query, err := r.repos.DB.Query()
	if err != nil {
		return &id, err
	}

	record := dbq.CreatePlanetParams{
		Name: args.Name,
	}
	if _, err := query.CreatePlanet(context.Background(), record); err != nil {
		return &id, nil
	}
	
	return &id, nil
}
