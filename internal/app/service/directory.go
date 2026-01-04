package service

import (
	"code/internal/pkg/errors"
	"os"
	"strings"
)

type Directory struct {
	Path               string
	IncludeHiddenFiles bool
}

func NewDirectory(path string, includeHiddenFiles bool) *Directory {
	return &Directory{Path: path, IncludeHiddenFiles: includeHiddenFiles}
}

func (dir Directory) GetSize() (int64, error) {
	var size int64 = 0

	files, err := os.ReadDir(dir.Path)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to read directory %s", dir.Path)
	}

	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return 0, errors.Wrapf(err, "failed to get file %s in %s", file, dir.Path)
		}

		if strings.HasPrefix(fileInfo.Name(), ".") && !dir.IncludeHiddenFiles {
			continue
		}

		if fileInfo.Mode().IsRegular() {
			size += fileInfo.Size()
		}
	}

	return size, nil
}
