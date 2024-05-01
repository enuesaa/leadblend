package service

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/enuesaa/leadblend/pkg/repository"
)

func NewPkgService(repos repository.Repos) PkgService {
	return PkgService{repos: repos}
}

type PkgService struct {
	repos repository.Repos
}

func (srv *PkgService) pkgFilename(name string) string {
	return fmt.Sprintf("%s.leadblend.zip", name)
}

func (srv *PkgService) pkgUnarchivedDBFilename(name string) string {
	return fmt.Sprintf(".leadblend/%s/data.db", name)
}

func (srv *PkgService) unarchivedir() string {
	return ".leadblend"
}

func (srv *PkgService) createUnarchiveDir() error {
	if srv.repos.Fs.IsExist(srv.unarchivedir()) {
		return nil
	}
	return srv.repos.Fs.CreateDir(srv.unarchivedir())
}

func (srv *PkgService) removeUnarchiveDir() error {
	return srv.repos.Fs.Remove(srv.unarchivedir())
}

func (srv *PkgService) IsExist(name string) bool {
	return srv.repos.Fs.IsExist(srv.pkgFilename(name))
}

func (srv *PkgService) archive(name string) (*bytes.Buffer, error) {
	b := bytes.NewBuffer([]byte{})
	writer := zip.NewWriter(b)
	defer writer.Close()

	f, err := os.Open(srv.pkgUnarchivedDBFilename(name))
	if err != nil {
		return b, err
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		return b, err
	}
	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return b, err
	}
	header.Name = "data.db"
	header.Method = zip.Deflate

	writerf, err := writer.CreateHeader(header)
	if err != nil {
		return b, err
	}
	if _, err = io.Copy(writerf, f); err != nil {
		return b, err
	}
	if err := writer.Flush(); err != nil {
		return b, err
	}
	return b, nil
}

func (srv *PkgService) Close(name string) error {
	b, err := srv.archive(name)
	if err != nil {
		return err
	}
	if err := srv.removeUnarchiveDir(); err != nil {
		return err
	}
	return srv.repos.Fs.Create(srv.pkgFilename(name), b.Bytes())
}

func (srv *PkgService) migrate(name string) error {
	path := srv.pkgUnarchivedDBFilename(name)
	if srv.repos.Fs.IsExist(path) {
		return nil
	}
	return srv.repos.DB.Migrate()
}

func (srv *PkgService) Create(name string) error {
	if err := srv.createUnarchiveDir(); err != nil {
		return err
	}
	if err := srv.repos.Fs.CreateDir(fmt.Sprintf(".leadblend/%s", name)); err != nil {
		return err
	}
	if err := srv.migrate(name); err != nil {
		return err
	}
	return srv.Close(name)
}

func (srv *PkgService) Open(name string) error {
	if err := srv.repos.Fs.CreateDir(".leadblend"); err != nil {
		return err
	}

	zr, err := zip.OpenReader(srv.pkgFilename(name))
	if err != nil {
		return err
	}
	defer zr.Close()

	if err := srv.repos.Fs.CreateDir(fmt.Sprintf(".leadblend/%s", name)); err != nil {
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

		path := filepath.Join(".leadblend", name, f.Name)
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
