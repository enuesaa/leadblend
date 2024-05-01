package usecase

import (
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/service"
)

func ClosePkg(repos repository.Repos, name string) error {
	pkgSrv := service.NewPkgService(repos)
	return pkgSrv.Close(name)
}
