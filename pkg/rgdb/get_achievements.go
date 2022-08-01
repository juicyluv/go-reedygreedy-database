package rgdb

import (
	"context"
	"fmt"
	"github.com/juicyluv/rgdb/pkg/rgdberr"
	"github.com/juicyluv/rgdb/pkg/rgdbmsg"
)

//language=PostgreSQL
const getAchievementsQuery = `
	select
	    achievement_id,
		name,
		description,
		payload,
		created_at,
		updated_at,
	    total
	from core.get_achievements(
	  _user_id := $1,
	  _page_size := $2,
	  _page := $3,
	  _sort := $4
	)
`

// TODO: fix user_id after db changes
func (c *Client) GetAchievements(ctx context.Context, request *rgdbmsg.GetAchievementsRequest) ([]*rgdbmsg.Achievement, int64, error) {
	request.CheckDefaults()

	rows, err := c.Driver.Query(
		ctx,
		getAchievementsQuery,

		request.UserId,
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

	err = rows.Scan(nil, nil, nil, nil, nil, nil, &total)

	if err != nil {
		return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
	}

	var achievements []*rgdbmsg.Achievement

	if int64(*request.PageSize) > total || *request.PageSize <= 0 {
		achievements = make([]*rgdbmsg.Achievement, 0, total)
	} else {
		achievements = make([]*rgdbmsg.Achievement, 0, *request.PageSize)
	}

	for rows.Next() {
		var achievement rgdbmsg.Achievement

		err = rows.Scan(
			&achievement.AchievementId,
			&achievement.Name,
			&achievement.Description,
			&achievement.Payload,
			&achievement.CreatedAt,
			&achievement.UpdatedAt,
			nil,
		)

		if err != nil {
			return nil, 0, fmt.Errorf(`%w: %v`, rgdberr.ErrInternal, err)
		}

		achievements = append(achievements, &achievement)
	}

	return achievements, total, nil
}
