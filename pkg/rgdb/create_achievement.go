package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const createAchievementQuery = `
	select 
		achievement_id,
		error
	from core.create_achievement(
	  _invoker_id := $1, 
	  _name := $2, 
	  _description := $3,
	  _payload := $4
	)
`

func (c *Client) CreateAchievement(ctx context.Context, request *rgdbmsg.CreateAchievementRequest) (*int64, error) {
	row, err := c.Driver.Query(
		ctx,
		createAchievementQuery,

		request.InvokerId,
		request.Name,
		request.Description,
		request.Payload,
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
		status        []byte
		achievementId *int64
	)

	err = row.Scan(
		&achievementId,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	if err = rgdberr.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	return achievementId, nil
}
