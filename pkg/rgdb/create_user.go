package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const createUserQuery = `
	select 
		user_id,
		error
	from core.create_user(
	  _invoker_id := $1, 
	  _username := $2, 
	  _email := $3, 
	  _password := $4, 
	  _timezone_id := $5,
	  _role_id := $6,
	  _avatar_url := $7,
	  _name := $8, 
	  _payload := $9
	)
`

func (c *Client) CreateUser(ctx context.Context, request *rgdbmsg.CreateUserRequest) (*int64, error) {
	row, err := c.Driver.Query(
		ctx,
		createUserQuery,

		request.InvokerId,
		request.Username,
		request.Email,
		request.Password,
		request.TimeZoneId,
		request.RoleId,
		request.AvatarURL,
		request.Name,
		request.Payload,
	)

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
		userId *int64
	)

	err = row.Scan(
		&userId,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	if err = rgdberr.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	return userId, nil
}
