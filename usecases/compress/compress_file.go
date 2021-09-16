package compress

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

type Archive struct {
	zw *zip.Writer
	f  *os.File
}

func CreateZip(filename string) (*Archive, error) {
	path := fmt.Sprintf("./tmp/compress/%s.zip", filename)
	f, err := os.Create(path)
	return &Archive{
		zw: zip.NewWriter(f),
		f:  f,
	}, err
}

func (a *Archive) Add(filename string, content []byte) error {
	var w io.Writer
	w, err := a.zw.Create(filename)
	if err != nil {
		return err
	}
	_, err = w.Write(content)
	return err
}

func (a *Archive) Close() error {
	err := a.zw.Close()
	if err != nil {
		return err
	}
	return a.f.Close()
}
