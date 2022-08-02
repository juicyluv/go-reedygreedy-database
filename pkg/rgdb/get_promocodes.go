package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const getPromocodesQuery = `
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
		total
	from core.get_promocodes(
	  _user_id := $1,
	  _search := $2,
	  _page_size := $3,
	  _page := $4,
	  _sort := $5
	)
`

func (c *Client) GetPromocodes(ctx context.Context, request *rgdbmsg.GetPromocodesRequest) ([]*rgdbmsg.Promocode, int64, error) {
	request.CheckDefaults()

	rows, err := c.Driver.Query(
		ctx,
		getPromocodesQuery,

		request.UserId,
		request.Search,
		request.PageSize,
		request.Page,
		request.Sort,
	)

	if err != nil {
		return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	defer rows.Close()

	if !rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		return nil, 0, rgdberr.ErrInternal
	}

	var total int64

	err = rows.Scan(nil, nil, nil, nil, nil, nil, nil, nil, nil, &total)

	if err != nil {
		return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	var promocodes []*rgdbmsg.Promocode

	if int64(*request.PageSize) > total || *request.PageSize <= 0 {
		promocodes = make([]*rgdbmsg.Promocode, 0, total)
	} else {
		promocodes = make([]*rgdbmsg.Promocode, 0, *request.PageSize)
	}

	for rows.Next() {
		var promocode rgdbmsg.Promocode

		err = rows.Scan(
			&promocode.PromocodeId,
			&promocode.Promocode,
			&promocode.Payload,
			&promocode.UsageCount,
			&promocode.CreatorId,
			&promocode.CreatorUsername,
			&promocode.CreatedAt,
			&promocode.UpdatedAt,
			&promocode.EndingAt,
			nil,
		)

		if err != nil {
			return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		promocodes = append(promocodes, &promocode)
	}

	return promocodes, total, nil
}
