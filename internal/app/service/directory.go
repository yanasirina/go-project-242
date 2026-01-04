package service

import (
	"code/internal/pkg/errors"
	"log"
	"os"
	"strings"
)

type Directory struct {
	Path               string
	Info               os.FileInfo
	IncludeHiddenFiles bool
}

func (dir Directory) GetSize() (int64, error) {
	var size int64 = 0

	files, err := os.ReadDir(dir.Path)
	if err != nil {
		log.Fatal(err)
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
