package rgdberrcode

const (
	PromocodeNotFound            ErrorCode = "PROMOCODE_NOT_FOUND"
	PromocodeAlreadyAddedToUser  ErrorCode = "PROMOCODE_ALREADY_ADDED_TO_USER"
	PromocodeAlreadyExists       ErrorCode = "PROMOCODE_ALREADY_EXISTS"
	PromocodeDoesNotBelongToUser ErrorCode = "PROMOCODE_DOES_NOT_BELONG_TO_USER"
)
