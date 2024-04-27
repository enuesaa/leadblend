package cli

import (
	"archive/zip"
	"io"
	"os"

	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateArchiveCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "archive",
		Short: "archive",
		RunE: func(cmd *cobra.Command, args []string) error {
			filename := "a.zip"
			fileToAdd := "README.md"

			reader, err := zip.OpenReader(filename)
			if err != nil {
				return err
			}
			defer reader.Close()
		
			existingZipFile, err := os.Create(filename)
			if err != nil {
				return err
			}
			defer existingZipFile.Close()
			writer := zip.NewWriter(existingZipFile)
			defer writer.Close()

			for _, existingFile := range reader.File {
				writerFile, err := writer.Create(existingFile.Name)
				if err != nil {
					return err
				}
				existingFileReader, err := existingFile.Open()
				if err != nil {
					return err
				}
				defer existingFileReader.Close()
				_, err = io.Copy(writerFile, existingFileReader)
				if err != nil {
					return err
				}
			}

			fileToAddReader, err := os.Open(fileToAdd)
			if err != nil {
				return err
			}
			defer fileToAddReader.Close()

			writerFile, err := writer.Create(fileToAdd)
			if err != nil {
				return err
			}
			_, err = io.Copy(writerFile, fileToAddReader)
			return err
		},
	}

	return cmd
}
