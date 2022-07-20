package rgdb

import (
	"context"
	"fmt"
	"rgdb/rgdberr"
	"rgdb/rgdbmsg"
)

//language=PostgreSQL
const getCategoriesQuery = `
	select
	    category_id,
	    name,
	    created_at,
	    updated_at,
	    total
	from core.get_categories(
	  _search := $1,
	  _page_size := $2,
	  _page := $3,
	  _sort := $4
	)
`

func (d *driver) GetCategories(ctx context.Context, request *rgdbmsg.GetCategoriesRequest) ([]*rgdbmsg.Category, int64, error) {
	request.CheckDefaults()

	rows, err := d.pool.Query(
		ctx,
		getCategoriesQuery,

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

	err = rows.Scan(nil, nil, nil, nil, &total)

	if err != nil {
		return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	var categories []*rgdbmsg.Category

	if int64(*request.PageSize) > total || *request.PageSize <= 0 {
		categories = make([]*rgdbmsg.Category, 0, total)
	} else {
		categories = make([]*rgdbmsg.Category, 0, *request.PageSize)
	}

	for rows.Next() {
		var category rgdbmsg.Category

		err = rows.Scan(
			&category.CategoryId,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
			nil,
		)

		if err != nil {
			return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		categories = append(categories, &category)
	}

	return categories, total, nil
}
