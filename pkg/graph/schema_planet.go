package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/graph-gophers/graphql-go"
)

type Planet struct {
	record dbq.Planet
}

func (p *Planet) Id() graphql.ID {
	return graphql.ID(p.record.ID)
}
func (p *Planet) Name() string {
	return p.record.Name
}
func (p *Planet) Comment() string {
	return p.record.Comment
}
