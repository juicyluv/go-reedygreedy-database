package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "rgdb/pkg/rgdberr"
	"rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const getBookReviewQuery = `
	select 
		book_id,
	    book_title,
	    creator_id,
	    creator_username,
	    review,
	    created_at,
	    updated_at,
	    comment,
		error
	from core.get_book_review(
	  _review_id := $1
	)
`

func (d *driver) GetBookReview(ctx context.Context, request *rgdbmsg.GetBookReviewRequest) (*rgdbmsg.BookReview, error) {
	row, err := d.pool.Query(ctx, getBookReviewQuery, request.ReviewId)

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
		status     []byte
		bookReview rgdbmsg.BookReview
	)

	err = row.Scan(
		&bookReview.BookId,
		&bookReview.BookTitle,
		&bookReview.CreatorId,
		&bookReview.CreatorUsername,
		&bookReview.Review,
		&bookReview.CreatedAt,
		&bookReview.UpdatedAt,
		&bookReview.Comment,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr2.ErrInternal, err)
	}

	if err = rgdberr2.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	bookReview.ReviewId = &request.ReviewId

	return &bookReview, nil
}
