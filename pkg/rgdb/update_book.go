package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const updateBookQuery = `
	select 
		update_book
	from core.update_book(
	  _invoker_id := $1, 
	  _book_id := $2,
	  _title := $3,
	  _price := $4,
	  _count := $5,
	  _author_id := $6,
	  _language_id := $7,
	  _description := $8
	)
`

func (d *driver) UpdateBook(ctx context.Context, request *rgdbmsg.UpdateBookRequest) error {
	row, err := d.pool.Query(
		ctx,
		updateBookQuery,

		request.InvokerId,
		request.BookId,
		request.Title,
		request.Price,
		request.Count,
		request.AuthorId,
		request.LanguageId,
		request.Description,
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
