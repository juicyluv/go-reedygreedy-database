package rgdberr_test

import (
	"errors"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdberr/rgdberrcode"
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
			name:          "request error",
			queryResponse: []byte(`{"status":1,"details":{"code":"USER_NOT_FOUND","message":"User not found.","error_type":"OBJECT_NOT_FOUND"}}`),
			expectedError: &rgdberr.DatabaseError{
				Code:      rgdberrcode.UserNotFound,
				Message:   "User not found.",
				ErrorType: rgdberr.ObjectNotFound,
			},
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
