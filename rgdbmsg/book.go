package rgdbmsg

import "time"

type (
	Book struct {
		BookId          *int64     `json:"book_id,omitempty"`
		Title           *string    `json:"title,omitempty"`
		Price           *float32   `json:"price,omitempty"`
		Count           *uint      `json:"count,omitempty"`
		CreatorId       *int64     `json:"creator_id,omitempty"`
		CreatorUsername *string    `json:"creator_username,omitempty"`
		AuthorId        *int64     `json:"author_id,omitempty"`
		AuthorName      *string    `json:"author_name,omitempty"`
		ISBN            *string    `json:"isbn,omitempty"`
		Pages           *uint16    `json:"pages,omitempty"`
		Language        *string    `json:"language,omitempty"`
		Description     *string    `json:"description,omitempty"`
		CreatedAt       *time.Time `json:"created_at,omitempty"`
		UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	}

	GetBookRequest struct {
		BookId int64
	}

	CreateBookRequest struct {
		InvokerId   int64
		Title       string
		Price       float32
		Count       int
		AuthorId    int64
		ISBN        string
		Pages       *int16
		LanguageId  *int16
		Description *string
	}

	GetBooksRequest struct {
		Search   *string
		PageSize *int
		Page     *int
		Sort     []string
	}

	UpdateBookRequest struct {
		InvokerId   int64
		BookId      int64
		Title       *string
		Price       *float32
		Count       *int
		AuthorId    *int64
		LanguageId  *int16
		Description *string
	}

	DeleteBookRequest struct {
		InvokerId int64
		BookId    int64
	}
)

func (req *GetBooksRequest) CheckDefaults() {
	_default(&req.PageSize, 60)
	_default(&req.Page, 1)
}
