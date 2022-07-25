package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const updateUserQuery = `
	select 
		update_user
	from core.update_user(
	  _invoker_id := $1, 
	  _user_id := $2, 
	  _username := $3,
	  _name := $4, 
	  _timezone_id := $5, 
	  _payload := $6
	)
`

func (c *Client) UpdateUser(ctx context.Context, request *rgdbmsg.UpdateUserRequest) error {
	row, err := c.Driver.Query(
		ctx,
		updateUserQuery,

		request.InvokerId,
		request.UserId,
		request.Username,
		request.Name,
		request.TimeZoneId,
		request.Payload,
	)

	if err != nil {
		return fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	defer row.Close()

	if !row.Next() {
		if err = row.Err(); err != nil {
			return fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		return rgdberr.ErrInternal
	}

	var status []byte

	err = row.Scan(&status)

	if err != nil {
		return fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	if err = rgdberr.AnalyzeQueryStatus(status); err != nil {
		return err
	}

	return nil
}
