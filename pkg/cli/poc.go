package cli

import (
	"archive/zip"
	"fmt"
	"io"

	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/spf13/cobra"
)

func CreatePocCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "poc",
		Short: "poc",
		RunE: func(cmd *cobra.Command, args []string) error {
			zr, err := zip.OpenReader("a.zip")
			if err != nil {
				return err
			}
			defer zr.Close()

			for _, f := range zr.File {
				reader, err := f.Open()
				if err != nil {
					return err
				}
				defer reader.Close()
				fbytes, err := io.ReadAll(reader)
				if err != nil {
					return err
				}
				fmt.Printf("%s", string(fbytes))
			}

			return nil
		},
	}

	return cmd
}
