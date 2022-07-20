package rgdbmsg

type (
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
