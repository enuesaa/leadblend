package graph

import (
	"github.com/graph-gophers/graphql-go"
)

type Island struct {
	id       string
	planetId string
	title    string
	content  string
	comment  string
}

func (p *Island) Id() graphql.ID {
	return graphql.ID(p.id)
}
func (p *Island) PlanetId() string {
	return p.planetId
}
func (p *Island) Title() string {
	return p.title
}
func (p *Island) Content() string {
	return p.content
}
func (p *Island) Comment() string {
	return p.comment
}
