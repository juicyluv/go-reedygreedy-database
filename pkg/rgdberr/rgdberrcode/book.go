package rgdberrcode

const (
	BookNotFound                 ErrorCode = "BOOK_NOT_FOUND"
	BookAlreadyAddedToCategory   ErrorCode = "BOOK_ALREADY_ADDED_TO_CATEGORY"
	BookAlreadyAddedToFavourites ErrorCode = "BOOK_ALREADY_ADDED_TO_FAVOURITES"
	BookIsNotInCategory          ErrorCode = "BOOK_IS_NOT_IN_CATEGORY"
	BookIsNotInFavourites        ErrorCode = "BOOK_NOT_IN_FAVOURITES"
)
