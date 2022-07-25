package rgdberrcode

const (
	UserEmailInvalidFormat   ErrorCode = "USER_EMAIL_INVALID_FORMAT"
	UserUsernameAlreadyTaken ErrorCode = "USERNAME_ALREADY_TAKEN"
	UserEmailAlreadyTaken    ErrorCode = "EMAIL_ALREADY_TAKEN"
	UserNotFound             ErrorCode = "USER_NOT_FOUND"
	UserNotDisabled          ErrorCode = "USER_NOT_DISABLED"
	UserAlreadyDisabled      ErrorCode = "USER_ALREADY_DISABLED"
)
