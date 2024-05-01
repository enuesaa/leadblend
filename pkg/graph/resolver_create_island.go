package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverCreateIslandArgs struct {
	PlanetId string
	Title    string
	Content  string
	Comment  string
}

func (r *Resolver) CreateIsland(args resolverCreateIslandArgs) (*string, error) {
	islandSrv := service.NewIslandService(r.repos)

	params := dbq.CreateIslandParams{
		PlanetID: args.PlanetId,
		Title:    args.Title,
		Content:  args.Content,
		Comment:  args.Comment,
	}
	id, err := islandSrv.Create(params)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
