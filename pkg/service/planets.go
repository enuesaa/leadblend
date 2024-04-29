package service

import "github.com/enuesaa/leadblend/pkg/repository"

type Planet struct{}

func NewPlanetService(repos repository.Repos) PlanetService {
	return PlanetService{
		repos: repos,
	}
}

type PlanetService struct {
	repos repository.Repos
}

func (srv *PlanetService) List() ([]Planet, error) {
	return make([]Planet, 0), nil
}
