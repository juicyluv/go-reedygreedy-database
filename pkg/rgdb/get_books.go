package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const getBooksQuery = `
	select
	    book_id,
		title,
	    count,
	    price,
		creator_id,
		creator_username,
		isbn,
		author_id,
		author_name,
		pages,
		language,
		description,
		created_at,
		updated_at,
	    total
	from core.get_books(
	  _search := $1,
	  _page_size := $2,
	  _page := $3,
	  _sort := $4
	)
`

func (c *Client) GetBooks(ctx context.Context, request *rgdbmsg.GetBooksRequest) ([]*rgdbmsg.Book, int64, error) {
	request.CheckDefaults()

	rows, err := c.pool.Query(
		ctx,
		getBooksQuery,

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

	err = rows.Scan(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &total)

	if err != nil {
		return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	var books []*rgdbmsg.Book

	if int64(*request.PageSize) > total || *request.PageSize <= 0 {
		books = make([]*rgdbmsg.Book, 0, total)
	} else {
		books = make([]*rgdbmsg.Book, 0, *request.PageSize)
	}

	for rows.Next() {
		var book rgdbmsg.Book

		err = rows.Scan(
			&book.BookId,
			&book.Title,
			&book.Count,
			&book.Price,
			&book.CreatorId,
			&book.CreatorUsername,
			&book.ISBN,
			&book.AuthorId,
			&book.AuthorName,
			&book.Pages,
			&book.Language,
			&book.Description,
			&book.CreatedAt,
			&book.UpdatedAt,
			nil,
		)

		if err != nil {
			return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		books = append(books, &book)
	}

	return books, total, nil
}
