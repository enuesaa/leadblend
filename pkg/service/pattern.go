package service

import (
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/oklog/ulid/v2"
)

func NewPatternService(repos repository.Repos) PatternService {
	return PatternService{
		repos: repos,
	}
}

type PatternService struct {
	repos repository.Repos
}

func (srv *PatternService) List() ([]dbq.Pattern, error) {
	query := srv.repos.DB.Query()
	return query.ListPatterns(ctx())
}

func (srv *PatternService) Get(id string) (dbq.Pattern, error) {
	query := srv.repos.DB.Query()
	return query.GetPattern(ctx(), id)
}

func (srv *PatternService) Create(params dbq.CreatePatternParams) (string, error) {
	query := srv.repos.DB.Query()
	params.ID = ulid.Make().String()
	if _, err := query.CreatePattern(ctx(), params); err != nil {
		return "", nil
	}
	return params.ID, nil
}
