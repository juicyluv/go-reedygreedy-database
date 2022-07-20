package rgdb

import (
	"context"
	"fmt"
	"rgdb/rgdberr"
	"rgdb/rgdbmsg"
)

//language=PostgreSQL
const deleteAuthorQuery = `
	select 
		delete_author
	from core.delete_author(
	  _invoker_id := $1, 
	  _author_id := $2
	)
`

func (d *driver) DeleteAuthor(ctx context.Context, request *rgdbmsg.DeleteAuthorRequest) error {
	row, err := d.pool.Query(
		ctx,
		deleteAuthorQuery,

		request.InvokerId,
		request.AuthorId,
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
