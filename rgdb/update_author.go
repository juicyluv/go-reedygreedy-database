package rgdb

import (
	"context"
	"fmt"
	"rgdb/rgdberr"
	"rgdb/rgdbmsg"
)

//language=PostgreSQL
const updateAuthorQuery = `
	select 
		update_author
	from core.update_author(
	  _invoker_id := $1, 
	  _author_id := $2,
	  _name := $3,
	  _description := $4
	)
`

func (d *driver) UpdateAuthor(ctx context.Context, request *rgdbmsg.UpdateAuthorRequest) error {
	row, err := d.pool.Query(
		ctx,
		updateAuthorQuery,

		request.InvokerId,
		request.AuthorId,
		request.Name,
		request.Description,
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
