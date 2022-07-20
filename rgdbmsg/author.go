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
)