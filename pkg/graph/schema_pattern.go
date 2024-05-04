package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/graph-gophers/graphql-go"
)

type Pattern struct {
	record dbq.Pattern
	traits []dbq.Trait
}

func (p *Pattern) Id() graphql.ID {
	return graphql.ID(p.record.ID)
}
func (p *Pattern) Title() string {
	return p.record.Title
}
func (p *Pattern) Traits() []*Trait {
	list := make([]*Trait, 0)
	for _, p := range p.traits {
		list = append(list, &Trait{p})
	}
	return list
}

type Trait struct {
	record dbq.Trait
}

func (p *Trait) Path() string {
	return p.record.Path
}
func (p *Trait) Type() string {
	return p.record.Type
}
func (p *Trait) DefaultValue() string {
	return p.record.DefaultValue
}
func (p *Trait) Required() bool {
	return p.record.Required
}
