package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const getAchievementQuery = `
	select 
		name,
		description,
		payload,
		created_at,
		updated_at,
		error
	from core.get_achievement(
	  _achievement_id := $1
	)
`

func (c *Client) GetAchievement(ctx context.Context, request *rgdbmsg.GetAchievementRequest) (*rgdbmsg.Achievement, error) {
	row, err := c.Driver.Query(ctx, getAchievementQuery, request.AchievementId)

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
		status      []byte
		achievement rgdbmsg.Achievement
	)

	err = row.Scan(
		&achievement.Name,
		&achievement.Description,
		&achievement.Payload,
		&achievement.CreatedAt,
		&achievement.UpdatedAt,
		&status,
	)

	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	if err = rgdberr.AnalyzeQueryStatus(status); err != nil {
		return nil, err
	}

	achievement.AchievementId = &request.AchievementId

	return &achievement, nil
}
