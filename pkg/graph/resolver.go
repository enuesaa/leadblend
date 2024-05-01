package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository"
)

type Resolver struct{
	repos repository.Repos
	db string
}
