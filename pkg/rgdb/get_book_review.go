package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
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

func (c *Client) GetBookReview(ctx context.Context, request *rgdbmsg.GetBookReviewRequest) (*rgdbmsg.BookReview, error) {
	row, err := c.Driver.Query(ctx, getBookReviewQuery, request.ReviewId)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	defer row.Close()

	if !row.Next() {
		if err = row.Err(); err != nil {
			return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		return nil, rgdberr.ErrInternal
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
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	if err = rgdberr.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	bookReview.ReviewId = &request.ReviewId

	return &bookReview, nil
}
