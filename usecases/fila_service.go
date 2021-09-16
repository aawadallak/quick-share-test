package usecases

import (
	"fmt"
	"log"
	"time"
	"we/domain"
	"we/dto"
	"we/infra/repository"
	"we/usecases/compress"
)

type FileService struct {
	repository domain.IOsFile
	repo       domain.IDatabase
}

func NewService(repository *repository.FileRepository, repo *repository.DatabaseRepository) FileService {
	return FileService{
		repository: repository,
		repo:       repo,
	}
}

func (f *FileService) UploadFile(files *dto.FileDTO) (interface{}, error) {

	upload := files.Convert2Entity()

	path, err := f.repository.CreateDirectory(upload)

	if err != nil {
		return "", err
	}

	_, err = f.repository.SaveOnDirectory(upload, path)

	if err != nil {
		return "", err
	}

	id, err := f.repo.Store(upload)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *FileService) DownloadFile(document string) (*string, error) {

	id, err := f.repo.FindById(document)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	files, err := f.repository.ReadAllFiles(id.GetPath())

	if err != nil {
		return nil, err
	}

	zip, err := compress.CreateZip(id.GetId())

	if err != nil {
		return nil, err
	}

	for _, file := range files {

		err := zip.Add(file.GetFileName(), file.GetContent())
		if err != nil {
			return nil, err
		}
	}

	err = zip.Close()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("./tmp/compress/%s.zip", id.GetId())

	return &path, err
}

func (f *FileService) CleanTempFiles() {

	for {

		documents, err := f.repo.FindAll()

		if err != nil {
			log.Println(err)
		}

		for _, d := range documents {
			err := f.repo.Delete()
			if err != nil {
				continue
			}
			err = f.repository.DeleteAllFilesFromFolder(d.GetPath())
			if err != nil {
				log.Println(err)
				continue
			}

			err = f.repository.FileExists(d.GetId())
			if err != nil {
				continue
			}

			err = f.repository.DeleteCompressFile(d.GetId())
			if err != nil {
				log.Println(err)
				continue
			}
		}

		time.Sleep(time.Minute * 1)
	}

}
