package rgdberr

import (
	"errors"
	"github.com/juicyluv/rgdb/pkg/rgdberr/rgdberrcode"
)

var (
	ErrInternal = errors.New(`internal error`)
	ErrUnknown  = errors.New(`unknown error`)
)

type ErrorType string

const (
	Unauthorized     ErrorType = "UNAUTHORIZED"
	InvalidArgument  ErrorType = "INVALID_ARGUMENT"
	ObjectNotFound   ErrorType = "OBJECT_NOT_FOUND"
	ObjectDuplicate  ErrorType = "OBJECT_DUPLICATE"
	ObjectDependency ErrorType = "OBJECT_DEPENDENCY"
)

type DatabaseError struct {
	Code      rgdberrcode.ErrorCode `json:"code"`
	Message   string                `json:"message"`
	ErrorType ErrorType             `json:"error_type"`
	Min       *float32              `json:"min,omitempty"`
	Max       *float32              `json:"max,omitempty"`
}

func (de *DatabaseError) Error() string {
	return de.Message
}
