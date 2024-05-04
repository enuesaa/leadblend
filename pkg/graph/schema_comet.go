package graph

import (
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/graph-gophers/graphql-go"
)

type Comet struct {
	record dbq.Stone
}

func (p *Comet) Id() graphql.ID {
	return graphql.ID(p.record.ID)
}
func (p *Comet) Data() string {
	return p.record.Data
}
