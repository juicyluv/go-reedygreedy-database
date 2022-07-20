package rgdb

import (
	"context"
	"fmt"
	"rgdb/rgdberr"
	"rgdb/rgdbmsg"
)

//language=PostgreSQL
const deleteCategoryQuery = `
	select 
		delete_category
	from core.delete_category(
	  _invoker_id := $1, 
	  _category_id := $2
	)
`

func (d *driver) DeleteCategory(ctx context.Context, request *rgdbmsg.DeleteCategoryRequest) error {
	row, err := d.pool.Query(
		ctx,
		deleteCategoryQuery,

		request.InvokerId,
		request.CategoryId,
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
