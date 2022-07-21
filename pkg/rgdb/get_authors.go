package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const getAuthorsQuery = `
	select
	    author_id,
	    name,
	    creator_id,
	    creator_username,
	    description,
	    created_at,
	    updated_at,
	    total
	from core.get_authors(
	  _search := $1,
	  _page_size := $2,
	  _page := $3,
	  _sort := $4
	)
`

func (d *driver) GetAuthors(ctx context.Context, request *rgdbmsg.GetAuthorsRequest) ([]*rgdbmsg.Author, int64, error) {
	request.CheckDefaults()

	rows, err := d.pool.Query(
		ctx,
		getAuthorsQuery,

		request.Search,
		request.PageSize,
		request.Page,
		request.Sort,
	)

	if err != nil {
		return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	defer rows.Close()

	if !rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		return nil, 0, rgdberr.ErrInternal
	}

	var total int64

	err = rows.Scan(nil, nil, nil, nil, nil, nil, nil, &total)

	if err != nil {
		return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	var authors []*rgdbmsg.Author

	if int64(*request.PageSize) > total || *request.PageSize <= 0 {
		authors = make([]*rgdbmsg.Author, 0, total)
	} else {
		authors = make([]*rgdbmsg.Author, 0, *request.PageSize)
	}

	for rows.Next() {
		var author rgdbmsg.Author

		err = rows.Scan(
			&author.AuthorId,
			&author.Name,
			&author.CreatorId,
			&author.CreatorUsername,
			&author.Description,
			&author.CreatedAt,
			&author.UpdatedAt,
			nil,
		)

		if err != nil {
			return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		authors = append(authors, &author)
	}

	return authors, total, nil
}
