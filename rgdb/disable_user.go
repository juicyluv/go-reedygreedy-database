package rgdb

import (
	"context"
	"fmt"
	"rgdb/rgdberr"
	"rgdb/rgdbmsg"
)

//language=PostgreSQL
const disableUserQuery = `
	select 
		disable_user
	from core.disable_user(
	  _invoker_id := $1, 
	  _user_id := $2, 
	  _disable_reason := $3
	)
`

func (d *driver) DisableUser(ctx context.Context, request *rgdbmsg.DisableUserRequest) error {
	row, err := d.pool.Query(
		ctx,
		disableUserQuery,

		request.InvokerId,
		request.UserId,
		request.DisableReason,
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
