package usecase

import (
	"strings"

	"github.com/enuesaa/leadblend/pkg/repository"
)

func FindLocalPkgs(repos repository.Repos) ([]string, error) {
	list := make([]string, 0)

	filenames, err := repos.Fs.ListFiles(".")
	if err != nil {
		return list, err
	}
	for _, filename := range filenames {
		if strings.HasSuffix(filename, ".leadblend.zip") {
			list = append(list, filename)
		}
	}

	return list, nil
}
