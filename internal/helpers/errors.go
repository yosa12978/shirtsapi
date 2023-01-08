package helpers

import "errors"

var (
	ErrNotFound   error = errors.New("entity not found")
	ErrBadRequest error = errors.New("bad request")
)
