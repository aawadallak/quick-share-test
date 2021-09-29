package dto

import (
	"mime/multipart"
	"quick_share/domain"
)

type FileDTO struct {
	File []*multipart.FileHeader
}

func (f *FileDTO) Convert2Entity() *domain.Upload {

	folder := domain.SetFolderPath()

	var file_domain []*domain.UploadFile

	for _, v := range f.File {
		dmn := domain.NewFile(v)
		file_domain = append(file_domain, dmn)
	}

	upload := domain.NewUpload(folder, file_domain)

	return upload
}
