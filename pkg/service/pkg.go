package service

import (
	"archive/zip"
	"bytes"
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

func (srv *PkgService) archive(name string) (*bytes.Buffer, error) {
	b := bytes.NewBuffer([]byte{})
	writer := zip.NewWriter(b)
	defer writer.Close()

	f, err := os.Open(".leadblend/data.db")
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

func (srv *PkgService) Create(name string) error {
	if err := srv.repos.Fs.CreateDir(".leadblend"); err != nil {
		return err
	}
	if err := srv.repos.DB.Migrate(".leadblend/data.db"); err != nil {
		return err
	}

	b, err := srv.archive(name)
	if err != nil {
		return err
	}
	srv.repos.Fs.Remove(".leadblend")
	return srv.repos.Fs.Create(srv.pkgFilename(name), b.Bytes())
}

func (srv *PkgService) RemoveOpened(name string) error {
	return srv.repos.Fs.Remove(srv.unarchivedir())
}

func (srv *PkgService) Open(name string) error {
	unarchivedir := fmt.Sprintf(".leadblend/%s", strings.ReplaceAll(name, ".zip", ""))

	zr, err := zip.OpenReader(name)
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
