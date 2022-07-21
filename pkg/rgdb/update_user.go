package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "github.com/juicyluv/rgdb/pkg/rgdberr"
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
	  _password := $4, 
	  _email := $5, 
	  _name := $6, 
	  _timezone_id := $7, 
	  _payload := $8
	)
`

func (d *driver) UpdateUser(ctx context.Context, request *rgdbmsg.UpdateUserRequest) error {
	row, err := d.pool.Query(
		ctx,
		updateUserQuery,

		request.InvokerId,
		request.UserId,
		request.Username,
		request.Password,
		request.Email,
		request.Name,
		request.TimeZoneId,
		request.Payload,
	)

	if err != nil {
		return fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
	}

	defer row.Close()

	if !row.Next() {
		if err = row.Err(); err != nil {
			return fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
		}

		return rgdberr2.ErrInternal
	}

	var status []byte

	err = row.Scan(&status)

	if err != nil {
		return fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
	}

	if err = rgdberr2.AnalyzeQueryStatus(status); err != nil {
		return err
	}

	return nil
}
