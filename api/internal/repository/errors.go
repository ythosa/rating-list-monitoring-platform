package repository

type DBError struct {
	err string
}

func NewDBError(err string) *DBError {
	return &DBError{err}
}

func (e *DBError) Error() string {
	return e.err
}

var (
	ErrRecordNotFound    = NewDBError("record not found")
	ErrUserAlreadyExists = NewDBError("user already exists")
)
