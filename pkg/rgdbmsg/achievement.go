package rgdbmsg

import "time"

type (
	Achievement struct {
		AchievementId *int16
		Name          *string
		Description   *string
		Payload       []byte
		CreatedAt     *time.Time
		UpdatedAt     *time.Time
	}

	GetAchievementRequest struct {
		AchievementId int16
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

	UpdateAchievementRequest struct {
		InvokerId     int64
		AchievementId int16
		Name          *string
		Description   *string
		Payload       []byte
	}

	DeleteAchievementRequest struct {
		InvokerId     int64
		AchievementId int16
	}
)

func (req *GetAchievementsRequest) CheckDefaults() {
	_default(&req.PageSize, 60)
	_default(&req.Page, 1)
}
