package service

import (
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/oklog/ulid/v2"
)

func NewStoneService(repos repository.Repos) StoneService {
	return StoneService{
		repos: repos,
	}
}

type StoneService struct {
	repos repository.Repos
}

func (srv *StoneService) List() ([]dbq.Stone, error) {
	query := srv.repos.DB.Query()
	return query.ListStones(ctx())
}

func (srv *StoneService) Get(id string) (dbq.Stone, error) {
	query := srv.repos.DB.Query()
	return query.GetStone(ctx(), id)
}

func (srv *StoneService) Create(params dbq.CreateStoneParams) (string, error) {
	query := srv.repos.DB.Query()
	// if _, err := query.GetIsland(ctx(), params.IslandID); err != nil {
	// 	return "", fmt.Errorf("planet id does not exist.")
	// }

	params.ID = ulid.Make().String()
	if _, err := query.CreateStone(ctx(), params); err != nil {
		return "", nil
	}
	return params.ID, nil
}

func (srv *StoneService) Delete(id string) error {
	query := srv.repos.DB.Query()
	return query.DeleteStone(ctx(), id)
}
