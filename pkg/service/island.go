package service

import (
	"context"
	"fmt"

	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/oklog/ulid/v2"
)

func NewIslandService(repos repository.Repos) IslandService {
	return IslandService{
		repos: repos,
	}
}

type IslandService struct {
	repos repository.Repos
}

func (srv *IslandService) List() ([]dbq.Island, error) {
	query, err := srv.repos.DB.Query()
	if err != nil {
		return make([]dbq.Island, 0), err
	}
	return query.ListIslands(context.Background())
}

func (srv *IslandService) Get(id string) (dbq.Island, error) {
	query, err := srv.repos.DB.Query()
	if err != nil {
		return dbq.Island{}, err
	}
	return query.GetIsland(context.Background(), id)
}

func (srv *IslandService) Create(params dbq.CreateIslandParams) (string, error) {
	query, err := srv.repos.DB.Query()
	if err != nil {
		return "", err
	}
	if _, err := query.GetPlanet(context.Background(), params.PlanetID); err != nil {
		return "", fmt.Errorf("planet id does not exist.")
	}

	params.ID = ulid.Make().String()
	if _, err := query.CreateIsland(context.Background(), params); err != nil {
		return "", nil
	}
	return params.ID, nil
}
