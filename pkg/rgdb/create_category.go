package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const createCategoryQuery = `
	select 
		category_id,
		error
	from core.create_category(
	  _invoker_id := $1, 
	  _name := $2
	)
`

func (c *Client) CreateCategory(ctx context.Context, request *rgdbmsg.CreateCategoryRequest) (*int64, error) {
	row, err := c.pool.Query(
		ctx,
		createCategoryQuery,

		request.InvokerId,
		request.Name,
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
		status     []byte
		categoryId *int64
	)

	err = row.Scan(
		&categoryId,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	if err = rgdberr.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	return categoryId, nil
}
