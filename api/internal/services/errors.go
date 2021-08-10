package services

type Error struct {
	Message string
}

func NewError(message string) *Error {
	return &Error{message}
}

func (e *Error) Error() string {
	return e.Message
}

var (
	UserAlreadyExistsError         = NewError("user already exists")
	InvalidUsernameOrPasswordError = NewError("invalid username or password")
	InvalidTokenError              = NewError("invalid token")
)
