package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const getBookQuery = `
	select 
		title,
	    count,
	    price,
	    creator_id,
	    creator_username,
	    author_id,
	    author_name,
	    isbn,
	    pages,
	    language,
	    description,
	    created_at,
	    updated_at,
		error
	from core.get_book(
	  _book_id := $1
	)
`

func (d *driver) GetBook(ctx context.Context, request *rgdbmsg.GetBookRequest) (*rgdbmsg.Book, error) {
	row, err := d.pool.Query(ctx, getBookQuery, request.BookId)

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
		book   rgdbmsg.Book
	)

	err = row.Scan(
		&book.Title,
		&book.Count,
		&book.Price,
		&book.CreatorId,
		&book.CreatorUsername,
		&book.AuthorId,
		&book.AuthorName,
		&book.ISBN,
		&book.Pages,
		&book.Language,
		&book.Description,
		&book.CreatedAt,
		&book.UpdatedAt,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
	}

	if err = rgdberr2.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	book.BookId = &request.BookId

	return &book, nil
}
