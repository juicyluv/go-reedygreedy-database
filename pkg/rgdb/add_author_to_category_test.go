package rgdb_test

import (
	"context"
	"errors"
	"github.com/juicyluv/rgdb/pkg/rgdb/rgdbtest"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
	"testing"
)

func TestClient_AddAuthorToCategory(t *testing.T) {
	const addAuthorToCategoryQuery = "select (.*) from core.add_author_to_category(.*)"

	var addAuthorToCategoryTestSet = []struct {
		name              string
		request           rgdbmsg.AddAuthorToCategoryRequest
		expectedBehaviour rgdbtest.ExpectedBehaviour
		expectedError     error
	}{
		{
			name: "Argument order",
			request: rgdbmsg.AddAuthorToCategoryRequest{
				InvokerId:  1,
				AuthorId:   2,
				CategoryId: 3,
			},
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query:   addAuthorToCategoryQuery,
				Args:    []interface{}{int64(1), int64(2), int16(3)},
				Columns: []string{"error"},
				Rows: [][]interface{}{
					{[]byte(`{"status":0}`)},
				},
				RowsError: nil,
			},
		},
		{
			name: "Check query status response",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query:   addAuthorToCategoryQuery,
				Columns: []string{"error"},
				Rows: [][]interface{}{
					{[]byte(`{"status":0}`)},
				},
				RowsError: nil,
			},
			expectedError: nil,
		},
		{
			name: "Request error",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query:   addAuthorToCategoryQuery,
				Columns: []string{"error"},
				Rows: [][]interface{}{
					{[]byte(`{"status":0}`)},
				},
				RowsError: errors.New("rows errors"),
			},
			expectedError: rgdberr.ErrInternal,
		},
		{
			name: "Unexpected row",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query:   addAuthorToCategoryQuery,
				Columns: []string{"error", "extra"},
				Rows: [][]interface{}{
					{[]byte(`{"status":0}`), "extra"},
				},
			},
			expectedError: rgdberr.ErrInternal,
		},
		{
			name: "Null response",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query: addAuthorToCategoryQuery,
			},
			expectedError: rgdberr.ErrInternal,
		},
	}

	for _, v := range addAuthorToCategoryTestSet {
		t.Run(v.name, func(t *testing.T) {
			client := rgdbtest.PrepareMock(t, v.expectedBehaviour)

			defer client.Close()

			err := client.AddAuthorToCategory(context.Background(), &v.request)

			rgdbtest.CheckResult(t, nil, nil, err, v.expectedError)
		})
	}
}
