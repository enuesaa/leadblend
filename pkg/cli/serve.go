package cli

import (
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			// pkgs, err := usecase.FindLocalPkgs(repos)
			// if err != nil {
			// 	return err
			// }
			// for _, pkg := range pkgs {
			// 	if err := usecase.OpenPkg(repos, pkg); err != nil {
			// 		return err
			// 	}
			// }

			if err := usecase.Migrate(repos); err != nil {
				return err
			}

			return usecase.Serve(repos)
		},
	}

	return cmd
}
