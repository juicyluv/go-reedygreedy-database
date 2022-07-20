package rgdb

import (
	"context"
	"fmt"
	"rgdb/rgdberr"
	"rgdb/rgdbmsg"
)

//language=PostgreSQL
const updateCategoryQuery = `
	select 
		update_category
	from core.update_category(
	  _invoker_id := $1,
	  _category_id := $2,
	  _name := $3
	)
`

func (d *driver) UpdateCategory(ctx context.Context, request *rgdbmsg.UpdateCategoryRequest) error {
	row, err := d.pool.Query(
		ctx,
		updateCategoryQuery,

		request.InvokerId,
		request.CategoryId,
		request.Name,
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
