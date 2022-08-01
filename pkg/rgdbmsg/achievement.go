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

	CreateAchievementRequest struct {
		InvokerId   int64
		Name        string
		Description string
		Payload     []byte
	}
)
