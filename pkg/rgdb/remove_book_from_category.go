package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const removeBookFromCategoryQuery = `
	select 
		remove_book_from_category
	from core.remove_book_from_category(
	  _invoker_id := $1, 
	  _book_id := $2,
	  _category_id := $3
	)
`

func (d *driver) RemoveBookFromCategory(ctx context.Context, request *rgdbmsg.RemoveBookFromCategoryRequest) error {
	row, err := d.pool.Query(
		ctx,
		removeBookFromCategoryQuery,

		request.InvokerId,
		request.BookId,
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
