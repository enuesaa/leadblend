package cli

import (
	"github.com/spf13/cobra"
)

func CreateServeCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
