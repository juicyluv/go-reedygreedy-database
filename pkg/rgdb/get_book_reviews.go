package rgdb

import (
	"context"
	"fmt"
	"rgdb/pkg/rgdberr"
	"rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const getBookReviewsQuery = `
	select
	    review_id,
	    book_id,
	    book_title,
	    creator_id,
	    creator_username,
	    review,
		created_at,
		updated_at,
	    comment,
	    total
	from core.get_book_reviews(
	  _book_id := $1,
	  _search := $2,
	  _page_size := $3,
	  _page := $4,
	  _sort := $5
	)
`

func (d *driver) GetBookReviews(ctx context.Context, request *rgdbmsg.GetBookReviewsRequest) ([]*rgdbmsg.BookReview, int64, error) {
	request.CheckDefaults()

	rows, err := d.pool.Query(
		ctx,
		getBookReviewsQuery,

		request.BookId,
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

	err = rows.Scan(nil, nil, nil, nil, nil, nil, nil, nil, nil, &total)

	if err != nil {
		return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	var bookReviews []*rgdbmsg.BookReview

	if int64(*request.PageSize) > total || *request.PageSize <= 0 {
		bookReviews = make([]*rgdbmsg.BookReview, 0, total)
	} else {
		bookReviews = make([]*rgdbmsg.BookReview, 0, *request.PageSize)
	}

	for rows.Next() {
		var bookReview rgdbmsg.BookReview

		err = rows.Scan(
			&bookReview.ReviewId,
			&bookReview.BookId,
			&bookReview.BookTitle,
			&bookReview.CreatorId,
			&bookReview.CreatorUsername,
			&bookReview.Review,
			&bookReview.CreatedAt,
			&bookReview.UpdatedAt,
			&bookReview.Comment,
			nil,
		)

		if err != nil {
			return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		bookReviews = append(bookReviews, &bookReview)
	}

	return bookReviews, total, nil
}
