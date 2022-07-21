package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "rgdb/pkg/rgdberr"
	"rgdb/pkg/rgdbmsg"
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

func (d *driver) CreateCategory(ctx context.Context, request *rgdbmsg.CreateCategoryRequest) (*int64, error) {
	row, err := d.pool.Query(
		ctx,
		createCategoryQuery,

		request.InvokerId,
		request.Name,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
	}

	defer row.Close()

	if !row.Next() {
		if err = row.Err(); err != nil {
			return nil, fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
		}

		return nil, rgdberr2.ErrInternal
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
		return nil, fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
	}

	if err = rgdberr2.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	return categoryId, nil
}
