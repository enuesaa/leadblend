package usecase

import (
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/service"
)

func OpenPkg(repos repository.Repos, filename string) error {
	pkgsrv := service.NewPkgService(repos)
	return pkgsrv.Open(filename)
}
