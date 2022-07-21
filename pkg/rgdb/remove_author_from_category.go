package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "rgdb/pkg/rgdberr"
	"rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const removeAuthorFromCategoryQuery = `
	select 
		remove_author_from_category
	from core.remove_author_from_category(
	  _invoker_id := $1, 
	  _author_id := $2,
	  _category_id := $3
	)
`

func (d *driver) RemoveAuthorFromCategory(ctx context.Context, request *rgdbmsg.RemoveAuthorFromCategoryRequest) error {
	row, err := d.pool.Query(
		ctx,
		removeAuthorFromCategoryQuery,

		request.InvokerId,
		request.AuthorId,
		request.CategoryId,
	)

	if err != nil {
		return fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
	}

	defer row.Close()

	if !row.Next() {
		if err = row.Err(); err != nil {
			return fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
		}

		return rgdberr2.ErrInternal
	}

	var status []byte

	err = row.Scan(&status)

	if err != nil {
		return fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
	}

	if err = rgdberr2.AnalyzeQueryStatus(status); err != nil {
		return err
	}

	return nil
}
