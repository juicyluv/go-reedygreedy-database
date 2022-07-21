package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
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
