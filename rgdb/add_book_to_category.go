package rgdb

import (
	"context"
	"fmt"
	"rgdb/rgdberr"
	"rgdb/rgdbmsg"
)

//language=PostgreSQL
const addBookToCategoryQuery = `
	select 
		add_book_to_category
	from core.add_book_to_category(
	  _invoker_id := $1, 
	  _book_id := $2,
	  _category_id := $3
	)
`

func (d *driver) AddBookToCategory(ctx context.Context, request *rgdbmsg.AddBookToCategoryRequest) error {
	row, err := d.pool.Query(
		ctx,
		addBookToCategoryQuery,

		request.InvokerId,
		request.BookId,
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
