package service

import (
	"context"

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
	query, err := srv.repos.DB.Query()
	if err != nil {
		return make([]dbq.Planet, 0), err
	}
	return query.ListPlanets(context.Background())
}

func (srv *PlanetService) Get(id string) (dbq.Planet, error) {
	query, err := srv.repos.DB.Query()
	if err != nil {
		return dbq.Planet{}, err
	}
	return query.GetPlanet(context.Background(), id)
}

func (srv *PlanetService) GetByName(name string) (dbq.Planet, error) {
	query, err := srv.repos.DB.Query()
	if err != nil {
		return dbq.Planet{}, err
	}
	return query.GetPlanetByName(context.Background(), name)
}

func (srv *PlanetService) Create(params dbq.CreatePlanetParams) (string, error) {
	query, err := srv.repos.DB.Query()
	if err != nil {
		return "", err
	}

	params.ID = ulid.Make().String()
	if _, err := query.CreatePlanet(context.Background(), params); err != nil {
		return "", nil
	}
	return params.ID, nil
}
