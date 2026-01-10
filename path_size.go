package code

import (
	"code/internal/app/handler"
	"fmt"
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
		return "", fmt.Errorf("get formated size of %s failed: %w", path, err)
	}

	return size, nil
}
