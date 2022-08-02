package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const deletePromocodeQuery = `
	select 
		delete_promocode
	from core.delete_promocode(
	  _invoker_id := $1, 
	  _promocode_id := $2
	)
`

func (c *Client) DeletePromocode(ctx context.Context, request *rgdbmsg.DeletePromocodeRequest) error {
	row, err := c.Driver.Query(
		ctx,
		deletePromocodeQuery,

		request.InvokerId,
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
