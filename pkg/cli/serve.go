package cli

import (
	"fmt"
	"strings"

	"github.com/enuesaa/leadblend/pkg/graph"
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			filename, _ := cmd.Flags().GetString("file")

			if filename != "" && !strings.HasSuffix(filename, ".leadblend.zip") {
				return fmt.Errorf("invalid filename.")
			}
			name := strings.ReplaceAll(filename, ".leadblend.zip", "")
			if err := usecase.UsePkg(repos, name); err != nil {
				return err
			}
			return graph.Serve(repos)
		},
	}
	cmd.Flags().String("file", "", "filename to open")

	return cmd
}
