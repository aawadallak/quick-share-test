package repository

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"quick_share/domain"
)

type FileRepository struct{}

func (o *FileRepository) CreateDirectory(upload *domain.Upload) (string, error) {

	path := fmt.Sprintf("./tmp/files/%s", upload.GetFolder().GetPath())

	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return path, nil
}

func (o *FileRepository) SaveOnDirectory(upload *domain.Upload, path string) (*domain.Upload, error) {

	files := upload.GetFiles()

	for _, f := range files {

		r := f.GetFile()

		// Open the file
		file, err := r.Open()
		if err != nil {
			return nil, err
		}

		defer file.Close()

		f, err := os.Create(fmt.Sprintf("./tmp/files/%s/%s", upload.GetFolder().GetPath(), f.GetFileName()))
		if err != nil {
			return nil, err
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			return nil, err
		}

	}

	return upload, nil
}

func (o *FileRepository) ReadAllFiles(id string) ([]*domain.File, error) {

	path := fmt.Sprintf("./tmp/files/%s", id)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirFiles []*domain.File

	for _, file := range files {

		pathFile := fmt.Sprintf("%s/%s", path, file.Name())

		content, err := ioutil.ReadFile(pathFile)

		if err != nil {
			return nil, err
		}

		f := domain.NewDownloadFile(file.Name(), file.IsDir(), file.Size(), content)
		dirFiles = append(dirFiles, f)
	}

	return dirFiles, nil
}

func (o *FileRepository) DeleteAllFilesFromFolder(path string) error {

	str := fmt.Sprintf("./tmp/files/%s", path)
	dir, err := ioutil.ReadDir(str)
	if err != nil {
		return err
	}

	for _, d := range dir {
		file := fmt.Sprintf("%s/%s", str, d.Name())
		err := os.RemoveAll(file)
		if err != nil {
			log.Println(err)
		}

	}

	err = os.Remove(str)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (o *FileRepository) DeleteCompressFile(file string) error {

	path := fmt.Sprintf("./tmp/compress/%s.zip", file)

	err := os.Remove(path)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (o *FileRepository) FileExists(file string) error {

	path := fmt.Sprintf("./tmp/compress/%s.zip", file)

	_, err := os.Stat(path)

	if err != nil {
		return err
	}

	return nil
}
