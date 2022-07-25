package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const deleteBookQuery = `
	select 
		delete_book
	from core.delete_book(
	  _invoker_id := $1, 
	  _book_id := $2
	)
`

func (c *Client) DeleteBook(ctx context.Context, request *rgdbmsg.DeleteBookRequest) error {
	row, err := c.Driver.Query(
		ctx,
		deleteBookQuery,

		request.InvokerId,
		request.BookId,
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
