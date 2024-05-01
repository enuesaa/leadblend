package service

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/enuesaa/leadblend/pkg/repository"
)


func NewPkgService(repos repository.Repos) PkgService {
	return PkgService{repos: repos}
}

type PkgService struct {
	repos repository.Repos
}

func (srv *PkgService) List() ([]string, error) {
	list := make([]string, 0)

	filenames, err := srv.repos.Fs.ListFiles(".")
	if err != nil {
		return list, err
	}
	for _, filename := range filenames {
		if strings.HasSuffix(filename, ".leadblend.zip") {
			list = append(list, filename)
		}
	}

	return list, nil
}

func (srv *PkgService) pkgFilename(name string) string {
	return fmt.Sprintf("%s.leadblend.zip", name)
}

func (srv *PkgService) unarchivedir() string {
	return ".leadblend"
}

func (srv *PkgService) IsExist(name string) bool {
	return srv.repos.Fs.IsExist(srv.pkgFilename(name))
}

func (srv *PkgService) Create(name string) error {
	// todo
	f, err := os.Create(srv.pkgFilename(name))
	if err != nil {
		return err
	}
	writer := zip.NewWriter(f)
	defer writer.Close()

	return writer.Flush()
}

func (srv *PkgService) RemoveOpened(name string) error {
	return srv.repos.Fs.Remove(srv.unarchivedir())
}

func (srv *PkgService) Open(filename string) error {
	unarchivedir := fmt.Sprintf(".leadblend/%s", strings.ReplaceAll(filename, ".zip", ""))

	zr, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer zr.Close()

	if err := srv.repos.Fs.CreateDir(unarchivedir); err != nil {
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
