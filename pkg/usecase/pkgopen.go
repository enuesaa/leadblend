package usecase

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/enuesaa/leadblend/pkg/repository"
)

func OpenPkg(repos repository.Repos, filename string) error {
	unarchivedir := fmt.Sprintf(".leadblend/%s", strings.ReplaceAll(filename, ".zip", ""))

	zr, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer zr.Close()

	if err := repos.Fs.CreateDir(unarchivedir); err != nil {
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

		path := filepath.Join(unarchivedir, filepath.Base(f.Name)) // remove parent dir name
		realwriter, err := os.Create(path)
		if err != nil {
			return err
		}
		defer realwriter.Close()
		if _, err := io.Copy(realwriter, reader); err != nil {
			return err
		}
	}

	return nil
}
