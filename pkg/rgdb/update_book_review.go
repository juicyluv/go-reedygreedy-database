package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "github.com/juicyluv/rgdb/pkg/rgdberr"
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

func (d *driver) UpdateBookReview(ctx context.Context, request *rgdbmsg.UpdateBookReviewRequest) error {
	row, err := d.pool.Query(
		ctx,
		updateBookReviewQuery,

		request.InvokerId,
		request.ReviewId,
		request.Review,
		request.Comment,
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
