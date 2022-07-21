package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const removeBookFromFavouritesQuery = `
	select 
		remove_book_from_favourites
	from core.remove_book_from_favourites(
	  _invoker_id := $1, 
	  _book_id := $2,
	  _user_id := $3
	)
`

func (c *Client) RemoveBookFromFavourites(ctx context.Context, request *rgdbmsg.RemoveBookFromFavouritesRequest) error {
	row, err := c.pool.Query(
		ctx,
		removeBookFromFavouritesQuery,

		request.InvokerId,
		request.BookId,
		request.UserId,
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
