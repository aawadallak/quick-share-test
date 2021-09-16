package domain

type File struct {
	fileName string
	isDir    bool
	size     int64
	content  []byte
}

func (d *File) GetFileName() string {
	return d.fileName
}

func (d *File) GetIsDir() bool {
	return d.isDir
}

func (d *File) GetFileSize() int64 {
	return d.size
}

func (d *File) GetContent() []byte {
	return d.content
}

func NewDownloadFile(filename string, isDir bool, size int64, content []byte) *File {
	return &File{
		fileName: filename,
		isDir:    isDir,
		size:     size,
		content:  content,
	}
}
