package rgdbmsg

import "time"

type (
	Author struct {
		AuthorId        *int64     `json:"author_id,omitempty"`
		Name            *string    `json:"name,omitempty"`
		CreatorId       *int64     `json:"creator_id,omitempty"`
		CreatorUsername *string    `json:"creator_username,omitempty"`
		Description     *string    `json:"description,omitempty"`
		CreatedAt       *time.Time `json:"created_at,omitempty"`
		UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	}

	GetAuthorRequest struct {
		AuthorId int64
	}

	GetAuthorsRequest struct {
		Search   *string
		PageSize *int
		Page     *int
		Sort     []string
	}

	CreateAuthorRequest struct {
		InvokerId   int64
		Name        string
		Description *string
	}

	UpdateAuthorRequest struct {
		InvokerId   int64
		AuthorId    int64
		Name        *string
		Description *string
	}

	DeleteAuthorRequest struct {
		InvokerId int64
		AuthorId  int64
	}

	AddAuthorToCategoryRequest struct {
		InvokerId  int64
		AuthorId   int64
		CategoryId int16
	}
)

func (req *GetAuthorsRequest) CheckDefaults() {
	_default(&req.PageSize, 60)
	_default(&req.Page, 1)
}
