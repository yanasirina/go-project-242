package handler

import "errors"

var ErrBadPath = errors.New("provided path is not a file or directory")
