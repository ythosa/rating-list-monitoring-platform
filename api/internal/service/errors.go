package service

type Error struct {
	Message string
}

func NewError(message string) *Error {
	return &Error{message}
}

func (s *Error) Error() string {
	return s.Message
}

var (
	UserAlreadyExistsError         = NewError("user already exists")
	InvalidUsernameOrPasswordError = NewError("invalid username or password")
	InvalidTokenError              = NewError("invalid token")
)
