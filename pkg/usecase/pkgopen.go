package usecase

import (
	"archive/zip"
	"fmt"
	"io"
)

func OpenPkg() error {
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
}
