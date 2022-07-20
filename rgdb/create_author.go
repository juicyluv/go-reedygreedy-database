package rgdb

import (
	"context"
	"fmt"
	"rgdb/rgdberr"
	"rgdb/rgdbmsg"
)

//language=PostgreSQL
const createAuthorQuery = `
	select 
		author_id,
		error
	from core.create_author(
	  _invoker_id := $1, 
	  _name := $2, 
	  _description := $3
	)
`

func (d *driver) CreateAuthor(ctx context.Context, request *rgdbmsg.CreateAuthorRequest) (*int64, error) {
	row, err := d.pool.Query(
		ctx,
		createAuthorQuery,

		request.InvokerId,
		request.Name,
		request.Description,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	defer row.Close()

	if !row.Next() {
		if err = row.Err(); err != nil {
			return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		return nil, rgdberr.ErrInternal
	}

	var (
		status   []byte
		authorId *int64
	)

	err = row.Scan(
		&authorId,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	if err = rgdberr.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	return authorId, nil
}