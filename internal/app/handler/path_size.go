package handler

import (
	"code/internal/app/service"
	"code/internal/pkg/errors"
	"code/internal/pkg/humanizer"
	"fmt"
	"os"
)

type Path interface {
	GetSize() (int64, error)
}

func (c PathSizeHandler) GetPath() (Path, error) {
	path := c.Arguments.Path
	includeHidden := c.Flags.ShowHiddenFiles

	pathInfo, err := os.Lstat(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to lstat %s", path)
	}

	switch mode := pathInfo.Mode(); {
	case mode.IsRegular():
		return service.NewFile(pathInfo), nil

	case mode.IsDir():
		return service.NewDirectory(path, includeHidden), nil
	}

	return nil, ErrBadPath
}

func (c PathSizeHandler) GetPathSize() (int64, error) {
	path := c.Arguments.Path

	pathInfo, err := c.GetPath()
	if err != nil {
		return 0, errors.Wrapf(err, "failed to get path %s", path)
	}

	size, err := pathInfo.GetSize()
	if err != nil {
		return 0, errors.Wrapf(err, "failed to get size of %s", path)
	}

	return size, nil
}

func (c PathSizeHandler) GetFormatedSize() (string, error) {
	size, err := c.GetPathSize()
	if err != nil {
		return "", errors.Wrap(err, "failed to get size")
	}

	if c.Flags.HumanizeSize {
		return humanizer.HumanizeBytes(size, 1000), nil
	} else {
		return fmt.Sprintf("%dB", size), nil
	}
}
