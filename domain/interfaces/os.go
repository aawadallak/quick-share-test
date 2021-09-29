package domain

import "quick_share/domain"

type IOsFile interface {
	CreateDirectory(*domain.Upload) (string, error)
	SaveOnDirectory(u *domain.Upload, path string) (*domain.Upload, error)
	ReadAllFiles(id string) ([]*domain.File, error)
	DeleteAllFilesFromFolder(path string) error
	DeleteCompressFile(file string) error
	FileExists(path string) error
}
