package usecase

import (
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/service"
)

func FindLocalPkgs(repos repository.Repos) ([]string, error) {
	pkgsrv := service.NewPkgService(repos)
	return pkgsrv.List()
}
