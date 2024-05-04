package service

import (
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
	query := srv.repos.DB.Query()
	return query.ListIslands(ctx())
}

func (srv *IslandService) ListByPlanetId(planetId string) ([]dbq.Island, error) {
	query := srv.repos.DB.Query()
	return query.ListIslandsByPlanetId(ctx(), planetId)
}

func (srv *IslandService) Get(id string) (dbq.Island, error) {
	query := srv.repos.DB.Query()
	return query.GetIsland(ctx(), id)
}

func (srv *IslandService) Create(params dbq.CreateIslandParams) (string, error) {
	query := srv.repos.DB.Query()
	if _, err := query.GetPlanet(ctx(), params.PlanetID); err != nil {
		return "", fmt.Errorf("planet id does not exist.")
	}

	params.ID = ulid.Make().String()
	if _, err := query.CreateIsland(ctx(), params); err != nil {
		return "", nil
	}
	return params.ID, nil
}

func (srv *IslandService) Delete(id string) error {
	query := srv.repos.DB.Query()
	return query.DeleteIsland(ctx(), id)
}
