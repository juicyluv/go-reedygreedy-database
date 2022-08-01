package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const updateAchievementQuery = `
	select 
		update_achievement
	from core.update_achievement(
	  _invoker_id := $1, 
	  _achievement_id := $2,
	  _name := $3,
	  _description := $4,
	  _payload := $5
	)
`

func (c *Client) UpdateAchievement(ctx context.Context, request *rgdbmsg.UpdateAchievementRequest) error {
	row, err := c.Driver.Query(
		ctx,
		updateAchievementQuery,

		request.InvokerId,
		request.AchievementId,
		request.Name,
		request.Description,
		request.Payload,
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
