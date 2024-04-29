package cli

import (
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/spf13/cobra"
	"net/http"
	_ "embed"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/99designs/gqlgen/graphql/playground"
)

//go:embed query.gql
var schema string

type Resolver struct{}
func (*Resolver) Planets() ([]Planet, error) {
	list := make([]Planet, 0)
	list = append(list, Planet{})
	return list, nil
}
func (*Resolver) CreatePlanet(args struct{ Name string }) (*string, error) {
	id := "a"
	return &id, nil
}

type Planet struct {}
func (p Planet) Id() graphql.ID {
	return graphql.ID("aaa")
}
func (p Planet) Name() string {
	return "aaaaaa"
}

func CreateGraphCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "graph",
		Short: "graph",
		RunE: func(cmd *cobra.Command, args []string) error {
			http.Handle("/graphql", &relay.Handler{
				Schema: graphql.MustParseSchema(schema, &Resolver{}),
			})
			http.Handle("/graphql/playground", playground.Handler("graph", "/graphql"))

			return http.ListenAndServe(":3000", nil)
		},
	}

	return cmd
}
