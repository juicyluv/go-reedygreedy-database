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

func TestClient_GetUsersTest(t *testing.T) {
	const getUsersQuery = "select (.*) from core.get_users(.*)"

	var (
		createdAt  = time.UnixMicro(1 << 10)
		updatedAt  = time.UnixMicro(1 << 20)
		disabledAt = time.UnixMicro(1 << 30)
		lastLogin  = time.UnixMicro(1 << 40)
	)

	var getUsersTestSet = []struct {
		name               string
		request            rgdbmsg.GetUsersRequest
		expectedBehaviour  rgdbtest.ExpectedBehaviour
		expectedError      error
		expectedResponse   []*rgdbmsg.User
		expectedTotalCount int64
	}{
		{
			name: "Request and response match",
			request: rgdbmsg.GetUsersRequest{
				Search:   ptr.String("search"),
				PageSize: ptr.Int(2),
				Page:     ptr.Int(3),
				Sort:     []string{"sort"},
			},
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query: getUsersQuery,
				Args: []interface{}{
					ptr.String("search"),
					ptr.Int(2),
					ptr.Int(3),
					[]string{"sort"},
				},
				Columns: []string{
					"user_id",
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
					"total",
				},
				Rows: [][]interface{}{
					{
						nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
						int64(2),
					},
					{
						ptr.Int64(1),
						ptr.String("username1"),
						ptr.String("email1"),
						nil,
						ptr.String("avatar_url1"),
						ptr.String("name1"),
						ptr.String("timezone1"),
						ptr.Int64(2),
						ptr.String("creator_username1"),
						ptr.Int16(3),
						ptr.String("role_name1"),
						ptr.Int16(4),
						&createdAt,
						&updatedAt,
						&disabledAt,
						ptr.Int16(5),
						&lastLogin,
						nil,
					},
					{
						ptr.Int64(2),
						ptr.String("username2"),
						ptr.String("email2"),
						nil,
						ptr.String("avatar_url2"),
						ptr.String("name2"),
						ptr.String("timezone2"),
						ptr.Int64(6),
						ptr.String("creator_username2"),
						ptr.Int16(7),
						ptr.String("role_name2"),
						ptr.Int16(8),
						&createdAt,
						&updatedAt,
						&disabledAt,
						ptr.Int16(9),
						&lastLogin,
						nil,
					},
				},
				RowsError: nil,
			},
			expectedResponse: []*rgdbmsg.User{
				{
					UserId:          ptr.Int64(1),
					Username:        ptr.String("username1"),
					Email:           ptr.String("email1"),
					Payload:         nil,
					Name:            ptr.String("name1"),
					AvatarURL:       ptr.String("avatar_url1"),
					TimeZone:        ptr.String("timezone1"),
					CreatorId:       ptr.Int64(2),
					CreatorUsername: ptr.String("creator_username1"),
					RoleId:          ptr.Int16(3),
					RoleName:        ptr.String("role_name1"),
					RoleAccessLevel: ptr.Int16(4),
					CreatedAt:       &createdAt,
					UpdatedAt:       &updatedAt,
					DisabledAt:      &disabledAt,
					DisableReason:   ptr.Int16(5),
					LastLogin:       &lastLogin,
				},
				{
					UserId:          ptr.Int64(2),
					Username:        ptr.String("username2"),
					Email:           ptr.String("email2"),
					Payload:         nil,
					Name:            ptr.String("name2"),
					AvatarURL:       ptr.String("avatar_url2"),
					TimeZone:        ptr.String("timezone2"),
					CreatorId:       ptr.Int64(6),
					CreatorUsername: ptr.String("creator_username2"),
					RoleId:          ptr.Int16(7),
					RoleName:        ptr.String("role_name2"),
					RoleAccessLevel: ptr.Int16(8),
					CreatedAt:       &createdAt,
					UpdatedAt:       &updatedAt,
					DisabledAt:      &disabledAt,
					DisableReason:   ptr.Int16(9),
					LastLogin:       &lastLogin,
				},
			},
			expectedTotalCount: 2,
		},
		{
			name: "Scan error",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query: getUsersQuery,
				Columns: []string{
					"user_id",
				},
				Rows: [][]interface{}{{
					ptr.Int64(1),
				}},
			},
			expectedError: rgdberr.ErrInternal,
		},
		{
			name: "Rows errors",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query:     getUsersQuery,
				RowsError: errors.New("rows error"),
			},
			expectedError: rgdberr.ErrInternal,
		},
		{
			name: "Unexpected column",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query:   getUsersQuery,
				Columns: []string{"user_id", "extra"},
				Rows: [][]interface{}{
					{ptr.Int64(1), "extra"},
				},
			},
			expectedError: rgdberr.ErrInternal,
		},
		{
			name: "Null response",
			expectedBehaviour: rgdbtest.ExpectedBehaviour{
				Query: getUsersQuery,
			},
			expectedError: rgdberr.ErrInternal,
		},
	}

	for _, v := range getUsersTestSet {
		t.Run(v.name, func(t *testing.T) {
			client := rgdbtest.PrepareMock(t, v.expectedBehaviour)

			defer client.Close()

			response, total, err := client.GetUsers(context.Background(), &v.request)

			if total != v.expectedTotalCount {
				t.Fatalf("Total count is not equal\nGot: %v\nExpected: %v\n", total, v.expectedTotalCount)
			}

			rgdbtest.CheckResult(t, response, v.expectedResponse, err, v.expectedError)
		})
	}
}
