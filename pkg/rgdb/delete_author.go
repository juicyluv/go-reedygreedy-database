package rgdb

import (
	"context"
	"fmt"
	rgdberr2 "github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
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
