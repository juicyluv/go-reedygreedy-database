package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const removePromocodeFromUserQuery = `
	select 
		remove_promocode_from_user
	from core.remove_promocode_from_user(
	  _invoker_id := $1, 
	  _user_id := $2,
	  _promocode_id := $3
	)
`

func (c *Client) RemovePromocodeFromUser(ctx context.Context, request *rgdbmsg.RemovePromocodeFromUserRequest) error {
	row, err := c.Driver.Query(
		ctx,
		removePromocodeFromUserQuery,

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
