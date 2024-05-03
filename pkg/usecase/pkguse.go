package usecase

import (
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/service"
)

func UsePkg(repos repository.Repos, filename string) error {
	pkgSrv := service.NewPkgService(repos)
	if !pkgSrv.IsExist(filename) {
		if err := pkgSrv.Create(filename); err != nil {
			return err
		}
	}
	return pkgSrv.Open(filename)
}
