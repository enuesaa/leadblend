package usecase

import (
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/service"
)

func PlanetsList(repos repository.Repos) ([]service.Planet, error) {
	planetSrv := service.NewPlanetService(repos)
	return planetSrv.List()
}
