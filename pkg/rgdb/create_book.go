package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const createBookQuery = `
	select 
		book_id,
		error
	from core.create_book(
	  _invoker_id := $1, 
	  _title := $2,
	  _price := $3,
	  _count := $4,
	  _author_id := $5,
	  _isbn := $6,
	  _pages := $7,
	  _language_id := $8,
	  _description := $9
	)
`

func (d *driver) CreateBook(ctx context.Context, request *rgdbmsg.CreateBookRequest) (*int64, error) {
	row, err := d.pool.Query(
		ctx,
		createBookQuery,

		request.InvokerId,
		request.Title,
		request.Price,
		request.Count,
		request.AuthorId,
		request.ISBN,
		request.Pages,
		request.LanguageId,
		request.Description,
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
		status []byte
		bookId *int64
	)

	err = row.Scan(
		&bookId,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
	}

	if err = rgdberr2.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	return bookId, nil
}
