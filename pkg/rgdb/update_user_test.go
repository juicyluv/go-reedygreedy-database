package rgdb_test

import (
	"context"
	"errors"
	"github.com/juicyluv/rgdb/pkg/rgdb/rgdbtest"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
	"github.com/juicyluv/rgutils/pkg/ptr"
	"testing"
)

func TestClient_UpdateUser(t *testing.T) {
	const updateUserQuery = "select (.*) from core.update_user(.*)"

	var updateUserTestSet = []struct {
		name              string
		request           rgdbmsg.UpdateUserRequest
		expectedBehaviour rgdbtest.ExpectedBehaviour
		expectedError     error
	}{
		{
			name: "Request and response match",
			request: rgdbmsg.UpdateUserRequest{
				InvokerId:  1,
				UserId:     2,
				Username:   ptr.String("username"),
				AvatarURL:  ptr.String("avatar_url"),
				Name:       ptr.String("name"),
				TimeZoneId: ptr.Int16(3),
				RoleId:     ptr.Int16(4),
				Payload:    []byte("payload"),
			},
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query: updateUserQuery,
				Args: []interface{}{
					int64(1),
					int64(2),
					ptr.String("username"),
					ptr.String("avatar_url"),
					ptr.String("name"),
					ptr.Int16(3),
					ptr.Int16(4),
					[]byte("payload"),
				},
				Columns: []string{
					"error",
				},
				Rows: [][]interface{}{{
					[]byte(`{"status": 0}`),
				}},
				RowsError: nil,
			},
		},
		{
			name: "Rows errors",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query:     updateUserQuery,
				RowsError: errors.New("rows error"),
			},
			expectedError: rgdberr.ErrInternal,
		},
		{
			name: "Unexpected column",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query:   updateUserQuery,
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
				Query: updateUserQuery,
			},
			expectedError: rgdberr.ErrInternal,
		},
		{
			name: "Database error",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query: updateUserQuery,
				Columns: []string{
					"error",
				},
				Rows: [][]interface{}{{
					[]byte(`{"status":1,"details":{"code": "something"}}`),
				}},
				RowsError: nil,
			},
			expectedError: &rgdberr.DatabaseError{
				Code: "something",
			},
		},
	}

	for _, v := range updateUserTestSet {
		t.Run(v.name, func(t *testing.T) {
			client := rgdbtest.PrepareMock(t, v.expectedBehaviour)

			defer client.Close()

			err := client.UpdateUser(context.Background(), &v.request)

			rgdbtest.CheckResult(t, nil, nil, err, v.expectedError)
		})
	}
}
