package code

import (
	"code/internal/app/handler"
	"code/internal/pkg/errors"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	pathSizeHandler := handler.PathSizeHandler{
		Arguments: handler.CommandArguments{
			Path: path,
		},
		Flags: handler.CommandFlags{
			HumanizeSize:    human,
			ShowHiddenFiles: all,
			Recursive:       recursive,
		},
	}

	size, err := pathSizeHandler.GetFormatedSize()
	if err != nil {
		return "", errors.Wrapf(err, "get formated size of %s failed", path)
	}

	return size, nil
}
