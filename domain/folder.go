package domain

import (
	"math/rand"
	"time"
)

type Folder struct {
	path string
}

func (f *Folder) GetPath() string {
	return f.path
}

func SetFolderPath() *Folder {
	str := generateHash()
	return &Folder{
		path: str,
	}
}

func generateHash() string {

	var seedRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[seedRand.Intn(len(charset))]
	}

	randomCode := string(b)

	return randomCode
}
