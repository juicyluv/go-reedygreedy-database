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

func TestClient_CreateUser(t *testing.T) {
	const createUserQuery = "select (.*) from core.create_user(.*)"

	var createUserTestSet = []struct {
		name              string
		request           rgdbmsg.CreateUserRequest
		expectedBehaviour rgdbtest.ExpectedBehaviour
		expectedError     error
		expectedResponse  *int64
	}{
		{
			name: "Request and response match",
			request: rgdbmsg.CreateUserRequest{
				InvokerId:  1,
				Username:   "username",
				Email:      "email",
				Password:   "password",
				TimeZoneId: 2,
				RoleId:     3,
				AvatarURL:  ptr.String("avatar_url"),
				Name:       ptr.String("Name"),
				Payload:    []byte(`payload`),
			},
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query: createUserQuery,
				Args: []interface{}{
					int64(1),
					"username",
					"email",
					"password",
					int16(2),
					int16(3),
					ptr.String("avatar_url"),
					ptr.String("Name"),
					[]byte(`payload`),
				},
				Columns: []string{
					"user_id",
					"error",
				},
				Rows: [][]interface{}{{
					ptr.Int64(1),
					[]byte(`{"status":0}`),
				}},
				RowsError: nil,
			},
			expectedResponse: ptr.Int64(1),
		},
		{
			name: "Scan error",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query: createUserQuery,
				Columns: []string{
					"error",
				},
				Rows: [][]interface{}{{
					[]byte(`{"status":0}`),
				}},
			},
			expectedError: rgdberr.ErrInternal,
		},
		{
			name: "Rows errors",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query:     createUserQuery,
				RowsError: errors.New("rows error"),
			},
			expectedError: rgdberr.ErrInternal,
		},
		{
			name: "Unexpected column",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query:   createUserQuery,
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
				Query: createUserQuery,
			},
			expectedError: rgdberr.ErrInternal,
		},
		{
			name: "Database error",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query: createUserQuery,
				Columns: []string{
					"user_id",
					"error",
				},
				Rows: [][]interface{}{{
					nil,
					[]byte(`{"status":1,"details":{"code": "something"}}`),
				}},
				RowsError: nil,
			},
			expectedError: &rgdberr.DatabaseError{
				Code: "something",
			},
		},
	}

	for _, v := range createUserTestSet {
		t.Run(v.name, func(t *testing.T) {
			client := rgdbtest.PrepareMock(t, v.expectedBehaviour)

			defer client.Close()

			response, err := client.CreateUser(context.Background(), &v.request)

			rgdbtest.CheckResult(t, response, v.expectedResponse, err, v.expectedError)
		})
	}
}
