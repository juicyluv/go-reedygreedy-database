package rgdbmsg

type (
	AddBookReviewRequest struct {
		InvokerId int64
		BookId    int64
		Review    int16
		Comment   *string
	}
)
