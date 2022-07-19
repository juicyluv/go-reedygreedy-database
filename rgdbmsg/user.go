package rgdbmsg

import "time"

type User struct {
	UserId          *int64     `json:"user_id,omitempty"`
	Username        *string    `json:"username,omitempty"`
	Email           *string    `json:"email,omitempty"`
	Payload         []byte     `json:"payload,omitempty"`
	Name            *string    `json:"name,omitempty"`
	TimeZone        *string    `json:"time_zone,omitempty"`
	CreatorId       *int64     `json:"creator_id,omitempty"`
	CreatorUsername *string    `json:"creator_username,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	DisabledAt      *time.Time `json:"disabled_at,omitempty"`
	DisableReason   *int16     `json:"disable_reason,omitempty"`
}

type GetUserRequest struct {
	UserId int64 `json:"user_id,omitempty"`
}
