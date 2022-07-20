package rgdb

import (
	"context"
	"fmt"
	"rgdb/rgdberr"
	"rgdb/rgdbmsg"
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
	  _name := $6, 
	  _payload := $7
	)
`

func (d *driver) CreateUser(ctx context.Context, request *rgdbmsg.CreateUserRequest) (*int64, error) {
	row, err := d.pool.Query(
		ctx,
		createUserQuery,

		request.InvokerId,
		request.Username,
		request.Email,
		request.Password,
		request.TimeZoneId,
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
