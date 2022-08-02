package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const getPromocodeByNameQuery = `
	select 
		promocode_id,
		promocode,
		payload,
		usage_count,
		creator_id,
		creator_username,
		created_at,
		updated_at,
		ending_at,
		error
	from core.get_promocode(
	  _promocode := $1
	)
`

func (c *Client) GetPromocodeByName(ctx context.Context, request *rgdbmsg.GetPromocodeByNameRequest) (*rgdbmsg.Promocode, error) {
	row, err := c.Driver.Query(ctx, getPromocodeByNameQuery, request.Promocode)

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
		status    []byte
		promocode rgdbmsg.Promocode
	)

	err = row.Scan(
		&promocode.PromocodeId,
		&promocode.Promocode,
		&promocode.Payload,
		&promocode.UsageCount,
		&promocode.CreatorId,
		&promocode.CreatorUsername,
		&promocode.CreatedAt,
		&promocode.UpdatedAt,
		&promocode.EndingAt,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	if err = rgdberr.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	return &promocode, nil
}
