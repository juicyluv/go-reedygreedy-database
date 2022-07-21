package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "rgdb/pkg/rgdberr"
	"rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const getCategoryQuery = `
	select 
		name,
		created_at,
	    updated_at,
		error
	from core.get_category(
	  _category_id := $1
	)
`

func (d *driver) GetCategory(ctx context.Context, request *rgdbmsg.GetCategoryRequest) (*rgdbmsg.Category, error) {
	row, err := d.pool.Query(ctx, getCategoryQuery, request.CategoryId)

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
		status   []byte
		category rgdbmsg.Category
	)

	err = row.Scan(
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
	}

	if err = rgdberr2.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	category.CategoryId = &request.CategoryId

	return &category, nil
}