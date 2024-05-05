package graph

import (
	"fmt"

	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverCreateCometArgs struct {
	Data string
}

func (r *Resolver) CreateComet(args resolverCreateCometArgs) (*string, error) {
	stoneSrv := service.NewStoneService(r.repos)
	patternSrv := service.NewPatternService(r.repos)

	params := dbq.CreateStoneParams{
		Data: args.Data,
	}
	id, err := stoneSrv.Create(params)
	if err != nil {
		return nil, err
	}
	if err := patternSrv.Evaluate(id); err != nil {
		// for debugging purpose.
		fmt.Printf("Error: %s\n", err.Error())
	}

	return &id, nil
}
