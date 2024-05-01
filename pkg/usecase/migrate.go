package usecase

import "github.com/enuesaa/leadblend/pkg/repository"

func Migrate(repos repository.Repos) error {
	if repos.DB.IsDBExist() {
		return nil
	}

	return repos.DB.Migrate() 
}
