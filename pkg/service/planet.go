package service

import (
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/oklog/ulid/v2"
)

func NewPlanetService(repos repository.Repos) PlanetService {
	return PlanetService{
		repos: repos,
	}
}

type PlanetService struct {
	repos repository.Repos
}

func (srv *PlanetService) List() ([]dbq.Planet, error) {
	query := srv.repos.DB.Query()
	return query.ListPlanets(ctx())
}

func (srv *PlanetService) Get(id string) (dbq.Planet, error) {
	query := srv.repos.DB.Query()
	return query.GetPlanet(ctx(), id)
}

func (srv *PlanetService) GetByName(name string) (dbq.Planet, error) {
	query := srv.repos.DB.Query()
	return query.GetPlanetByName(ctx(), name)
}

func (srv *PlanetService) Create(params dbq.CreatePlanetParams) (string, error) {
	query := srv.repos.DB.Query()
	params.ID = ulid.Make().String()
	if _, err := query.CreatePlanet(ctx(), params); err != nil {
		return "", nil
	}
	return params.ID, nil
}
