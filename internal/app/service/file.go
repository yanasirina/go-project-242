package service

import "os"

type File struct {
	Path string
	Info os.FileInfo
}

func (file File) GetSize() (int64, error) {
	return file.Info.Size(), nil
}
