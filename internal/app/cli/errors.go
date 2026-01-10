package cli

import "errors"

var ErrBadArguments = errors.New("command expects exactly one argument")
