package graph

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type Planet struct{
	id string
	name string
}

func (p *Planet) Id() graphql.ID {
	return graphql.ID(p.id)
}
func (p *Planet) Name() string {
	return p.name
}
