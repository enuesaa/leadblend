package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverCreatePatternArgs struct {
	Title string
}

func (r *Resolver) CreatePattern(args resolverCreatePatternArgs) (*string, error) {
	patternSrv := service.NewPatternService(r.repos)

	params := dbq.CreatePatternParams{
		Title:    args.Title,
	}
	id, err := patternSrv.Create(params)
	if err != nil {
		return nil, err
	}

	patternSrv.AddTrait(dbq.CreateTraitParams{
		PatternID: id,
		Path: "$.a",
		Type: "string",
		DefaultValue: "a",
		Required: true,
	})
	patternSrv.AddTrait(dbq.CreateTraitParams{
		PatternID: id,
		Path: "$.b",
		Type: "string",
		DefaultValue: "b",
		Required: true,
	})

	return &id, nil
}
