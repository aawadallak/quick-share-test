package domain

import "mime/multipart"

type UploadFile struct {
	info *multipart.FileHeader
}

func (f *UploadFile) GetFileName() string {
	return f.info.Filename
}

func (f *UploadFile) GetFileSize() int {
	return int(f.info.Size)
}

func (f *UploadFile) GetFile() *multipart.FileHeader {
	return f.info
}

func NewFile(file *multipart.FileHeader) *UploadFile {
	return &UploadFile{
		info: file,
	}
}
