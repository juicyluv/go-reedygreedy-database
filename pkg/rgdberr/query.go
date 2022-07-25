package rgdberr

import (
	"encoding/json"
	"fmt"
)

const (
	queryStatusInternal = iota - 1
	queryStatusSuccess
	queryStatusFailure
)

type queryStatus struct {
	Status  int8           `json:"status"`
	Details *DatabaseError `json:"details,omitempty"`
}

func AnalyzeQueryStatus(status []byte) error {
	if status == nil {
		return fmt.Errorf("%w: no query status provided", ErrInternal)
	}

	var query queryStatus

	err := json.Unmarshal(status, &query)

	if err != nil {
		return fmt.Errorf(`%w: %v`, ErrInternal, err)
	}

	switch query.Status {
	case queryStatusSuccess:
		return nil
	case queryStatusInternal:
		return ErrUnknown
	case queryStatusFailure:
		if query.Details == nil {
			return fmt.Errorf("%w: no error details", ErrInternal)
		}

		return query.Details
	default:
		return fmt.Errorf("%w: unhandled query status", ErrInternal)
	}
}
