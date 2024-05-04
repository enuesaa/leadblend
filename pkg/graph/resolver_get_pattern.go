package graph

import "github.com/enuesaa/leadblend/pkg/service"

type resolverPatternArgs struct {
	Id *string
}

func (r *Resolver) GetPattern(args resolverPatternArgs) (*Pattern, error) {
	patternSrv := service.NewPatternService(r.repos)
	pattern, err := patternSrv.GetWithTrait(*args.Id)
	if err != nil {
		return nil, err
	}
	res := Pattern{
		record: pattern.Pattern,
		traits: pattern.Traits,
	}
	return &res, nil
}
