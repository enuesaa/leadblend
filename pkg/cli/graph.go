package cli

import (
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/spf13/cobra"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (query) Hello() string {
	return "Hello, world!"
}

func CreateGraphCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "graph",
		Short: "graph",
		RunE: func(cmd *cobra.Command, args []string) error {
			s := `
				type Query {
						hello: String!
				}
			`
			schema := graphql.MustParseSchema(s, &query{})
			http.Handle("/query", &relay.Handler{Schema: schema})

			return http.ListenAndServe(":3000", nil)
		},
	}

	return cmd
}
