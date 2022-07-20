package rgdbmsg

import "time"

type (
	BookReview struct {
		ReviewId        *int64     `json:"review_id,omitempty"`
		BookId          *int64     `json:"book_id,omitempty"`
		BookTitle       *string    `json:"book_title,omitempty"`
		CreatorId       *int64     `json:"creator_id,omitempty"`
		CreatorUsername *string    `json:"creator_username,omitempty"`
		Review          *int16     `json:"review,omitempty"`
		CreatedAt       *time.Time `json:"created_at,omitempty"`
		UpdatedAt       *time.Time `json:"updated_at,omitempty"`
		Comment         *string    `json:"comment,omitempty"`
	}

	GetBookReviewsRequest struct {
		BookId   *int64
		Search   *string
		PageSize *int
		Page     *int
		Sort     []string
	}

	AddBookReviewRequest struct {
		InvokerId int64
		BookId    int64
		Review    int16
		Comment   *string
	}

	UpdateBookReviewRequest struct {
		InvokerId int64
		ReviewId  int64
		Review    *int16
		Comment   *string
	}

	RemoveBookReviewRequest struct {
		InvokerId int64
		ReviewId  int64
	}
)

func (req *GetBookReviewsRequest) CheckDefaults() {
	_default(&req.PageSize, 60)
	_default(&req.Page, 1)
}
