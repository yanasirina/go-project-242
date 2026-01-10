package handler

import (
	"fmt"
	"os"

	"code/internal/app/service"
	"code/internal/pkg/humanizer"
)

type Path interface {
	GetSize() (int64, error)
}

func (c PathSizeHandler) GetPath() (Path, error) {
	path := c.Arguments.Path
	includeHidden := c.Flags.ShowHiddenFiles
	recursive := c.Flags.Recursive

	pathInfo, err := os.Lstat(path)
	if err != nil {
		return nil, fmt.Errorf("failed to lstat %s: %w", path, err)
	}

	mode := pathInfo.Mode()
	if mode.IsDir() {
		return service.NewDirectory(path, includeHidden, recursive), nil
	}

	return service.NewFile(pathInfo), nil
}

func (c PathSizeHandler) GetPathSize() (int64, error) {
	path := c.Arguments.Path

	pathInfo, err := c.GetPath()
	if err != nil {
		return 0, fmt.Errorf("failed to get path %s: %w", path, err)
	}

	size, err := pathInfo.GetSize()
	if err != nil {
		return 0, fmt.Errorf("failed to get size of %s: %w", path, err)
	}

	return size, nil
}

func (c PathSizeHandler) GetFormatedSize() (string, error) {
	size, err := c.GetPathSize()
	if err != nil {
		return "", fmt.Errorf("failed to get size: %w", err)
	}

	if c.Flags.HumanizeSize {
		return humanizer.HumanizeBytes(size, 1024), nil
	} else {
		return fmt.Sprintf("%dB", size), nil
	}
}
