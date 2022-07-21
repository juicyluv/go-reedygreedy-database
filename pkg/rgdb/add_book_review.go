package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "rgdb/pkg/rgdberr"
	"rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const addBookReviewQuery = `
	select 
		review_id,
		error
	from core.add_book_review(
	  _invoker_id := $1, 
	  _book_id := $2,
	  _review := $3,
	  _comment := $4
	)
`

func (d *driver) AddBookReview(ctx context.Context, request *rgdbmsg.AddBookReviewRequest) (*int64, error) {
	row, err := d.pool.Query(
		ctx,
		addBookReviewQuery,

		request.InvokerId,
		request.BookId,
		request.Review,
		request.Comment,
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
		status   []byte
		reviewId *int64
	)

	err = row.Scan(
		&reviewId,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
	}

	if err = rgdberr2.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	return reviewId, nil
}
