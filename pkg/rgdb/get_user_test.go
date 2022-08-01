package rgdb_test

import (
	"context"
	"errors"
	"github.com/juicyluv/rgdb/pkg/rgdb/rgdbtest"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
	"github.com/juicyluv/rgutils/pkg/ptr"
	"testing"
	"time"
)

func TestClient_GetUser(t *testing.T) {
	const getUserQuery = "select (.*) from core.get_user(.*)"

	var (
		createdAt  = time.UnixMicro(1 << 10)
		updatedAt  = time.UnixMicro(1 << 20)
		disabledAt = time.UnixMicro(1 << 30)
		lastLogin  = time.UnixMicro(1 << 40)
	)

	var getUserTestSet = []struct {
		name              string
		request           rgdbmsg.GetUserRequest
		expectedBehaviour rgdbtest.ExpectedBehaviour
		expectedError     error
		expectedResponse  *rgdbmsg.User
	}{
		{
			name: "Request and response match",
			request: rgdbmsg.GetUserRequest{
				UserId: 1,
			},
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query: getUserQuery,
				Args:  []interface{}{int64(1)},
				Columns: []string{
					"username",
					"email",
					"payload",
					"avatar_url",
					"name",
					"timezone",
					"creator_id",
					"creator_username",
					"role_id",
					"role_name",
					"role_access_level",
					"created_at",
					"updated_at",
					"disabled_at",
					"disable_reason",
					"last_login",
					"error",
				},
				Rows: [][]interface{}{{
					ptr.String("username"),
					ptr.String("email"),
					nil,
					ptr.String("avatar_url"),
					ptr.String("name"),
					ptr.String("timezone"),
					ptr.Int64(2),
					ptr.String("creator_username"),
					ptr.Int16(3),
					ptr.String("role_name"),
					ptr.Int16(4),
					&createdAt,
					&updatedAt,
					&disabledAt,
					ptr.Int16(5),
					&lastLogin,
					[]byte(`{"status":0}`),
				}},
				RowsError: nil,
			},
			expectedResponse: &rgdbmsg.User{
				UserId:          ptr.Int64(1),
				Username:        ptr.String("username"),
				Email:           ptr.String("email"),
				Payload:         nil,
				AvatarURL:       ptr.String("avatar_url"),
				Name:            ptr.String("name"),
				TimeZone:        ptr.String("timezone"),
				CreatorId:       ptr.Int64(2),
				CreatorUsername: ptr.String("creator_username"),
				RoleId:          ptr.Int16(3),
				RoleName:        ptr.String("role_name"),
				RoleAccessLevel: ptr.Int16(4),
				CreatedAt:       &createdAt,
				UpdatedAt:       &updatedAt,
				DisabledAt:      &disabledAt,
				DisableReason:   ptr.Int16(5),
				LastLogin:       &lastLogin,
			},
		},
		{
			name: "Scan error",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query: getUserQuery,
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
				Query:     getUserQuery,
				RowsError: errors.New("rows error"),
			},
			expectedError: rgdberr.ErrInternal,
		},
		{
			name: "Unexpected column",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query:   getUserQuery,
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
				Query: getUserQuery,
			},
			expectedError: rgdberr.ErrInternal,
		},
		{
			name: "Database error",
			request: rgdbmsg.GetUserRequest{
				UserId: 1,
			},
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query: getUserQuery,
				Args:  []interface{}{int64(1)},
				Columns: []string{
					"username",
					"email",
					"payload",
					"avatar_url",
					"name",
					"timezone",
					"creator_id",
					"creator_username",
					"role_id",
					"role_name",
					"role_access_level",
					"created_at",
					"updated_at",
					"disabled_at",
					"disable_reason",
					"last_login",
					"error",
				},
				Rows: [][]interface{}{{
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
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

	for _, v := range getUserTestSet {
		t.Run(v.name, func(t *testing.T) {
			client := rgdbtest.PrepareMock(t, v.expectedBehaviour)

			defer client.Close()

			response, err := client.GetUser(context.Background(), &v.request)

			rgdbtest.CheckResult(t, response, v.expectedResponse, err, v.expectedError)
		})
	}
}
