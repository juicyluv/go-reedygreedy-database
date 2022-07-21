package rgdberr

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	queryStatusInternal = iota - 1
	queryStatusSuccess
	queryStatusFailure
)

type queryStatus struct {
	Status  int8          `json:"status"`
	Details *queryDetails `json:"details,omitempty"`
}

type queryDetails struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

func AnalyzeQueryStatus(status []byte) error {
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
			return ErrInternal
		}

		return errors.New(query.Details.Message)
	default:
		return ErrUnknown
	}
}
