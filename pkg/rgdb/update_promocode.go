package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const updatePromocodeQuery = `
	select 
		update_promocode
	from core.update_promocode(
	  _invoker_id := $1, 
	  _promocode_id := $2,
	  _promocode := $3,
	  _payload := $4,
	  _usage_count := $5,
	  _ending_at := $6
	)
`

func (c *Client) UpdatePromocode(ctx context.Context, request *rgdbmsg.UpdatePromocodeRequest) error {
	row, err := c.Driver.Query(
		ctx,
		updatePromocodeQuery,

		request.InvokerId,
		request.PromocodeId,
		request.Promocode,
		request.Payload,
		request.UsageCount,
		request.EndingAt,
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
