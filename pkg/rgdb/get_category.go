package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
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
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	if err = rgdberr.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	category.CategoryId = &request.CategoryId

	return &category, nil
}
