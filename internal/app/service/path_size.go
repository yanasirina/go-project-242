package service

import (
	"code/internal/pkg/errors"
	"os"
)

type Path interface {
	GetSize() (int64, error)
}

func GetPath(path string) (Path, error) {
	pathInfo, err := os.Lstat(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to lstat %s", path)
	}

	switch mode := pathInfo.Mode(); {
	case mode.IsRegular():
		return File{Path: path, Info: pathInfo}, nil

	case mode.IsDir():
		return Directory{Path: path, Info: pathInfo}, nil
	}

	return nil, ErrBadPath
}

func GetPathSize(path string) (int64, error) {
	pathInfo, err := GetPath(path)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to get path %s", path)
	}

	size, err := pathInfo.GetSize()
	if err != nil {
		return 0, errors.Wrapf(err, "failed to get size of %s", path)
	}

	return size, nil
}
