package domain

type IOsFile interface {
	CreateDirectory(*Upload) (string, error)
	SaveOnDirectory(u *Upload, path string) (*Upload, error)
	ReadAllFiles(id string) ([]*File, error)
	DeleteAllFilesFromFolder(path string) error
	DeleteCompressFile(file string) error
	FileExists(path string) error
}
