package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const addPromocodeToUserQuery = `
	select 
		add_promocode_to_user
	from core.add_promocode_to_user(
	  _invoker_id := $1, 
	  _user_id := $2,
	  _promocode_id := $3
	)
`

func (c *Client) AddPromocodeToUser(ctx context.Context, request *rgdbmsg.AddPromocodeToUserRequest) error {
	row, err := c.Driver.Query(
		ctx,
		addPromocodeToUserQuery,

		request.InvokerId,
		request.UserId,
		request.PromocodeId,
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
