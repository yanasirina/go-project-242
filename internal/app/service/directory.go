package service

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/yanasirina/go-project-242/internal/pkg/errors"
)

type Directory struct {
	Path               string
	IncludeHiddenFiles bool
	Recursive          bool
}

func NewDirectory(path string, includeHiddenFiles, recursive bool) *Directory {
	return &Directory{Path: path, IncludeHiddenFiles: includeHiddenFiles, Recursive: recursive}
}

func (dir Directory) GetSize() (int64, error) {
	var folderSize int64 = 0

	files, err := os.ReadDir(dir.Path)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to read directory %s", dir.Path)
	}

	for _, file := range files {
		size, err := dir.fileSize(file)
		if err != nil {
			return 0, err
		}

		folderSize += size
	}

	return folderSize, nil
}

func (dir Directory) fileSize(file os.DirEntry) (int64, error) {
	fileInfo, err := file.Info()
	if err != nil {
		return 0, errors.Wrapf(err, "failed to get file %s in %s", file, dir.Path)
	}

	if strings.HasPrefix(fileInfo.Name(), ".") && !dir.IncludeHiddenFiles {
		return 0, nil
	}

	if fileInfo.Mode().IsRegular() {
		return fileInfo.Size(), nil
	}

	if dir.Recursive && fileInfo.IsDir() {
		recDirPath := filepath.Join(dir.Path, fileInfo.Name())
		recDir := NewDirectory(recDirPath, dir.IncludeHiddenFiles, dir.Recursive)

		return recDir.GetSize()
	}

	return 0, nil
}
