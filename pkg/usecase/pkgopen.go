package usecase

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/enuesaa/leadblend/pkg/repository"
)

func OpenPkg(repos repository.Repos) error {
	zr, err := zip.OpenReader("a.zip")
	if err != nil {
		return err
	}
	defer zr.Close()

	if err := repos.Fs.CreateDir(".a"); err != nil {
		return err
	}

	for _, f := range zr.File {
		reader, err := f.Open()
		if err != nil {
			return err
		}
		defer reader.Close()

		if f.FileInfo().IsDir() {
			continue
		}

		path := filepath.Join(".a", filepath.Base(f.Name)) // remove parent dir name
		realwriter, err := os.Create(path)
		if err != nil {
			return err
		}
		defer realwriter.Close()
		if _,  err := io.Copy(realwriter, reader); err != nil {
			return err
		}
	}

	return nil
}
