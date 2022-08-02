package rgdbmsg

import "time"

type (
	Promocode struct {
		PromocodeId     *int64     `json:"promocode_id,omitempty"`
		Promocode       *string    `json:"promocode,omitempty"`
		Payload         []byte     `json:"payload,omitempty"`
		UsageCount      *int       `json:"usage_count,omitempty"`
		CreatorId       *int64     `json:"creator_id,omitempty"`
		CreatorUsername *string    `json:"creator_username,omitempty"`
		CreatedAt       *time.Time `json:"created_at,omitempty"`
		UpdatedAt       *time.Time `json:"updated_at,omitempty"`
		EndingAt        *time.Time `json:"ending_at,omitempty"`
	}

	CreatePromocodeRequest struct {
		InvokerId  int64
		Promocode  string
		Payload    []byte
		UsageCount *int
		EndingAt   *time.Time
	}

	GetPromocodesRequest struct {
		UserId   *int64
		Search   *string
		PageSize *int
		Page     *int
		Sort     []string
	}

	DeletePromocodeRequest struct {
		InvokerId   int64
		PromocodeId int64
	}

	AddPromocodeToUserRequest struct {
		InvokerId   int64
		UserId      int64
		PromocodeId int64
	}
)

func (req *GetPromocodesRequest) CheckDefaults() {
	_default(&req.PageSize, 60)
	_default(&req.Page, 1)
}
