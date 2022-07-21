package rgdberr

import "errors"

var (
	ErrInternal = errors.New(`internal error`)
	ErrUnknown  = errors.New(`unknown error`)
)
