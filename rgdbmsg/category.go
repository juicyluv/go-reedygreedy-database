package rgdbmsg

import "time"

type (
	Category struct {
		CategoryId *int16     `json:"category_id,omitempty"`
		Name       *string    `json:"name,omitempty"`
		CreatedAt  *time.Time `json:"created_at,omitempty"`
		UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	}

	GetCategoryRequest struct {
		CategoryId int16
	}

	CreateCategoryRequest struct {
		InvokerId int64
		Name      string
	}

	GetCategoriesRequest struct {
		Search   *string
		PageSize *int
		Page     *int
		Sort     []string
	}
)

func (req *GetCategoriesRequest) CheckDefaults() {
	_default(&req.PageSize, 60)
	_default(&req.Page, 1)
}
