package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/graph-gophers/graphql-go"
)

type Stone struct {
	record dbq.Stone
}

func (p *Stone) Id() graphql.ID {
	return graphql.ID(p.record.ID)
}
func (p *Stone) PatternId() *string {
	if p.record.PatternID.String == "" {
		return nil
	}
	return &p.record.PatternID.String
}
func (p *Stone) IslandId() *string {
	if p.record.IslandID.String == "" {
		return nil
	}
	return &p.record.IslandID.String
}
func (p *Stone) Data() string {
	return p.record.Data
}
