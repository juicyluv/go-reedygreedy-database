package rgdbmsg

import "time"

type (
	Category struct {
		CategoryId *int16     `json:"category_id,omitempty"`
		Name       *string    `json:"name,omitempty"`
		CreatedAt  *time.Time `json:"created_at,omitempty"`
	}

	GetCategoryRequest struct {
		CategoryId int16
	}
)
