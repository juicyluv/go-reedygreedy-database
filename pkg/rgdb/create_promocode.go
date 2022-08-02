package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const createPromocodeQuery = `
	select 
		promocode_id,
		error
	from core.create_promocode(
	  _invoker_id := $1, 
	  _promocode := $2, 
	  _payload := $3,
	  _usage_count := $4,
	  _ending_at := $5
	)
`

func (c *Client) CreatePromocode(ctx context.Context, request *rgdbmsg.CreatePromocodeRequest) (*int64, error) {
	row, err := c.Driver.Query(
		ctx,
		createPromocodeQuery,

		request.InvokerId,
		request.Promocode,
		request.Payload,
		request.UsageCount,
		request.EndingAt,
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
		status      []byte
		promocodeId *int64
	)

	err = row.Scan(
		&promocodeId,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	if err = rgdberr.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	return promocodeId, nil
}
