package graph

import (
	"context"

	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/oklog/ulid/v2"
)

type CreatePlanetArgs struct {
	Name string
}
func (r *Resolver) CreatePlanet(args CreatePlanetArgs) (*string, error) {
	if err := r.repos.DB.Open(r.db); err != nil {
		return nil, err
	}

	query, err := r.repos.DB.Query()
	if err != nil {
		return nil, err
	}

	id := ulid.Make().String()
	record := dbq.CreatePlanetParams{
		ID: id,
		Name: args.Name,
	}
	if _, err := query.CreatePlanet(context.Background(), record); err != nil {
		return &id, nil
	}
	
	return &id, nil
}
