package rgdbmsg

import "time"

type (
	User struct {
		UserId          *int64     `json:"user_id,omitempty"`
		Username        *string    `json:"username,omitempty"`
		Email           *string    `json:"email,omitempty"`
		Payload         []byte     `json:"payload,omitempty"`
		Name            *string    `json:"name,omitempty"`
		AvatarURL       *string    `json:"avatar_url,omitempty"`
		TimeZone        *string    `json:"time_zone,omitempty"`
		CreatorId       *int64     `json:"creator_id,omitempty"`
		CreatorUsername *string    `json:"creator_username,omitempty"`
		RoleId          *int16     `json:"role_id,omitempty"`
		RoleName        *string    `json:"role_name,omitempty"`
		RoleAccessLevel *int16     `json:"role_access_level,omitempty"`
		CreatedAt       *time.Time `json:"created_at,omitempty"`
		UpdatedAt       *time.Time `json:"updated_at,omitempty"`
		DisabledAt      *time.Time `json:"disabled_at,omitempty"`
		DisableReason   *int16     `json:"disable_reason,omitempty"`
		LastLogin       *time.Time `json:"last_login,omitempty"`
	}

	GetUserRequest struct {
		UserId int64
	}

	GetUsersRequest struct {
		Search   *string
		PageSize *int
		Page     *int
		Sort     []string
	}

	CreateUserRequest struct {
		InvokerId  int64
		Username   string
		Email      string
		Password   string
		TimeZoneId int16
		RoleId     int16
		AvatarURL  *string
		Name       *string
		Payload    []byte
	}

	UpdateUserRequest struct {
		InvokerId  int64
		UserId     int64
		Username   *string
		AvatarURL  *string
		Name       *string
		TimeZoneId *int16
		RoleId     *int16
		Payload    []byte
	}

	DisableUserRequest struct {
		InvokerId     int64
		UserId        int64
		DisableReason int16
	}
)

func (req *GetUsersRequest) CheckDefaults() {
	_default(&req.PageSize, 60)
	_default(&req.Page, 1)
}
