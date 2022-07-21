package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const updateBookReviewQuery = `
	select 
		update_book_review
	from core.update_book_review(
	  _invoker_id := $1,
	  _review_id := $2,
	  _review := $3,
	  _comment := $4
	)
`

func (c *Client) UpdateBookReview(ctx context.Context, request *rgdbmsg.UpdateBookReviewRequest) error {
	row, err := c.pool.Query(
		ctx,
		updateBookReviewQuery,

		request.InvokerId,
		request.ReviewId,
		request.Review,
		request.Comment,
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
