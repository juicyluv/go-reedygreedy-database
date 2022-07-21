package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "rgdb/pkg/rgdberr"
	"rgdb/pkg/rgdbmsg"
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

func (d *driver) RemoveBookFromFavourites(ctx context.Context, request *rgdbmsg.RemoveBookFromFavouritesRequest) error {
	row, err := d.pool.Query(
		ctx,
		removeBookFromFavouritesQuery,

		request.InvokerId,
		request.BookId,
		request.UserId,
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
