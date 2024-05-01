package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/graph-gophers/graphql-go"
)

type Island struct {
	record dbq.Island
}

func (p *Island) Id() graphql.ID {
	return graphql.ID(p.record.ID)
}
func (p *Island) PlanetId() string {
	return p.record.PlanetID
}
func (p *Island) Title() string {
	return p.record.Title
}
func (p *Island) Content() string {
	return p.record.Content
}
func (p *Island) Comment() string {
	return p.record.Comment
}
