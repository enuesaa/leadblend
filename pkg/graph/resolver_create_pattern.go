package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/enuesaa/leadblend/pkg/service"
)

type resolverCreatePatternArgs struct {
	Title string
	Traits []resolverCreatePatternTrait
}
type resolverCreatePatternTrait struct {
	Path string
	Type string
	DefaultValue string
	Required bool
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

	for _, t := range args.Traits {
		patternSrv.AddTrait(dbq.CreateTraitParams{
			PatternID: id,
			Path: t.Path,
			Type: t.Type,
			DefaultValue: t.DefaultValue,
			Required: t.Required,
		})
	}

	return &id, nil
}
