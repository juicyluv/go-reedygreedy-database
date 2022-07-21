package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const getUsersQuery = `
	select
	    user_id,
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
	    total
	from core.get_users(
	  _search := $1,
	  _page_size := $2,
	  _page := $3,
	  _sort := $4
	)
`

func (c *Client) GetUsers(ctx context.Context, request *rgdbmsg.GetUsersRequest) ([]*rgdbmsg.User, int64, error) {
	request.CheckDefaults()

	rows, err := c.pool.Query(
		ctx,
		getUsersQuery,

		request.Search,
		request.PageSize,
		request.Page,
		request.Sort,
	)

	if err != nil {
		return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	defer rows.Close()

	if !rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		return nil, 0, rgdberr.ErrInternal
	}

	var total int64

	err = rows.Scan(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &total)

	if err != nil {
		return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	var users []*rgdbmsg.User

	if int64(*request.PageSize) > total || *request.PageSize <= 0 {
		users = make([]*rgdbmsg.User, 0, total)
	} else {
		users = make([]*rgdbmsg.User, 0, *request.PageSize)
	}

	for rows.Next() {
		var user rgdbmsg.User

		err = rows.Scan(
			&user.UserId,
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
			nil,
		)

		if err != nil {
			return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		users = append(users, &user)
	}

	return users, total, nil
}
