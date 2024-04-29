package usecase

import (
	"github.com/enuesaa/leadblend/pkg/graph"
	"github.com/enuesaa/leadblend/pkg/repository"
)

func Serve(repos repository.Repos) error {
	return graph.Serve(repos)
}
