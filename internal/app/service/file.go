package service

import "os"

type File struct {
	Info os.FileInfo
}

func NewFile(info os.FileInfo) *File {
	return &File{Info: info}
}

func (file File) GetSize() (int64, error) {
	return file.Info.Size(), nil
}
