package rgdberr_test

import (
	"errors"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdberr/rgdberrcode"
	"github.com/juicyluv/rgutils/pkg/ptr"
	"reflect"
	"testing"
)

func TestAnalyzeQueryStatus(t *testing.T) {
	testTable := []struct {
		name          string
		queryResponse []byte
		expectedError error
	}{
		{
			name:          "valid response",
			queryResponse: []byte(`{"status":0}`),
		},
		{
			name:          "unknown error",
			queryResponse: []byte(`{"status":-1}`),
			expectedError: rgdberr.ErrUnknown,
		},
		{
			name: "request error",
			queryResponse: []byte(
				`{
					"status": 1,
					"details": {
						"code": "USER_NOT_FOUND",
						"message": "User not found.",
						"error_type": "OBJECT_NOT_FOUND"
					}
				 }`,
			),
			expectedError: &rgdberr.DatabaseError{
				Code:      rgdberrcode.UserNotFound,
				Message:   "User not found.",
				ErrorType: rgdberr.ObjectNotFound,
			},
		},
		{
			name: "value out of range",
			queryResponse: []byte(
				`{
					"status": 1,
					"details": {
						"code": "VALUE_OUT_OF_RANGE",
						"message": "Author name is out of range.",
						"error_type": "INVALID_ARGUMENT",
						"min": 4,
						"max": 100
					}
				 }`,
			),
			expectedError: &rgdberr.DatabaseError{
				Code:      rgdberrcode.ValueOutOfRange,
				Message:   "Author name is out of range.",
				ErrorType: rgdberr.InvalidArgument,
				Min:       ptr.Float32(4),
				Max:       ptr.Float32(100),
			},
		},
		{
			name:          "empty object",
			queryResponse: []byte(`{}`),
			expectedError: rgdberr.ErrInternal,
		},
		{
			name:          "no status field",
			queryResponse: []byte(`{"details": {"something": 123}}`),
			expectedError: rgdberr.ErrInternal,
		},
		{
			name:          "empty json string",
			queryResponse: []byte(``),
			expectedError: rgdberr.ErrInternal,
		},
		{
			name:          "no details field",
			queryResponse: []byte(`{"status": 1}`),
			expectedError: rgdberr.ErrInternal,
		},
		{
			name:          "invalid json",
			queryResponse: []byte(`{"status"}`),
			expectedError: rgdberr.ErrInternal,
		},
		{
			name:          "invalid status",
			queryResponse: []byte(`{status:1000"}`),
			expectedError: rgdberr.ErrInternal,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			err := rgdberr.AnalyzeQueryStatus(tt.queryResponse)

			if !errors.Is(err, tt.expectedError) {
				if !reflect.DeepEqual(err, tt.expectedError) {
					t.Fatalf("Mismatch error\n\nGot:%+v\nExpected:%+v", err, tt.expectedError)
				}
			}
		})
	}
}
