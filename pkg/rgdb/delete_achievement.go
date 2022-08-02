package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const deleteAchievementQuery = `
	select 
		delete_achievement
	from core.delete_achievement(
	  _invoker_id := $1, 
	  _achievement_id := $2
	)
`

func (c *Client) DeleteAchievement(ctx context.Context, request *rgdbmsg.DeleteAchievementRequest) error {
	row, err := c.Driver.Query(
		ctx,
		deleteAchievementQuery,

		request.InvokerId,
		request.AchievementId,
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
