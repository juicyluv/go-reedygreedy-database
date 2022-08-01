package rgdbmsg

import "time"

type (
	Achievement struct {
		AchievementId *int64
		Name          *string
		Description   *string
		Payload       []byte
		CreatedAt     *time.Time
		UpdatedAt     *time.Time
	}

	GetAchievementRequest struct {
		AchievementId int64
	}

	GetAchievementsRequest struct {
		UserId   *int64
		PageSize *int
		Page     *int
		Sort     []string
	}

	CreateAchievementRequest struct {
		InvokerId   int64
		Name        string
		Description string
		Payload     []byte
	}
)

func (req *GetAchievementsRequest) CheckDefaults() {
	_default(&req.PageSize, 60)
	_default(&req.Page, 1)
}
