package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const addBookToFavouritesQuery = `
	select 
		add_book_to_favourites
	from core.add_book_to_favourites(
	  _invoker_id := $1, 
	  _book_id := $2,
	  _user_id := $3
	)
`

func (d *driver) AddBookToFavourites(ctx context.Context, request *rgdbmsg.AddBookToFavouritesRequest) error {
	row, err := d.pool.Query(
		ctx,
		addBookToFavouritesQuery,

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
