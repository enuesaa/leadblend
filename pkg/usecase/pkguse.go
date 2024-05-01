package usecase

import (
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/service"
)

func UsePkg(repos repository.Repos, name string) error {
	if name == "" {
		name = "main"
	}

	pkgSrv := service.NewPkgService(repos)
	if !pkgSrv.IsExist(name) {
		if err := pkgSrv.Create(name); err != nil {
			return err
		}
	}

	// if err := pkgSrv.RemoveOpened(name); err != nil {
	// 	return err
	// }

	// if err := pkgSrv.Open(name); err != nil {
	// 	return err
	// }

	return nil
}