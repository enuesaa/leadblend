package graph

import "github.com/enuesaa/leadblend/pkg/service"

func (r *Resolver) ListPatterns() ([]*Pattern, error) {
	list := make([]*Pattern, 0)

	patternSrv := service.NewPatternService(r.repos)
	patterns, err := patternSrv.List()
	if err != nil {
		return list, err
	}
	for _, p := range patterns {
		list = append(list, &Pattern{
			record: p,
		})
	}

	return list, nil
}
