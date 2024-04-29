package graph

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type Planet struct{}

func (p Planet) Id() graphql.ID {
	return graphql.ID("aaa")
}
func (p Planet) Name() string {
	return "aaaaaa"
}
