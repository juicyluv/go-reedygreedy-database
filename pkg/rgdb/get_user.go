package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const getUserQuery = `
	select 
		username,
	    email,
	    payload,
		name,
		timezone,
		creator_id,
		creator_username,
		created_at,
		updated_at,
		disabled_at,
		disable_reason,
		error
	from core.get_user(
	  _user_id := $1
	)
`

func (c *Client) GetUser(ctx context.Context, request *rgdbmsg.GetUserRequest) (*rgdbmsg.User, error) {
	row, err := c.Driver.Query(ctx, getUserQuery, request.UserId)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	defer row.Close()

	if !row.Next() {
		if err = row.Err(); err != nil {
			return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		return nil, rgdberr.ErrInternal
	}

	var (
		status []byte
		user   rgdbmsg.User
	)

	err = row.Scan(
		&user.Username,
		&user.Email,
		&user.Payload,
		&user.Name,
		&user.TimeZone,
		&user.CreatorId,
		&user.CreatorUsername,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DisabledAt,
		&user.DisableReason,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	if err = rgdberr.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	user.UserId = &request.UserId

	return &user, nil
}
