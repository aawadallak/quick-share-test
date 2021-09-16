package domain

type Upload struct {
	folder *Folder
	files  []*UploadFile
}

func (u *Upload) GetFolder() *Folder {
	return u.folder
}

func (u *Upload) GetFiles() []*UploadFile {
	return u.files
}

func NewUpload(f *Folder, file []*UploadFile) *Upload {
	return &Upload{
		folder: f,
		files:  file,
	}
}
