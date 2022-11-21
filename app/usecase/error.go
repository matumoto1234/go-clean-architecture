package usecase

const (
	// ErrInternalServerError : Internal Server Error
	ErrInternalServerError = iota + 1
	// ErrNotFound : Not Found Error
	ErrNotFound
	// ErrBadRequest : Bad Request Error
	ErrBadRequest
)

// Error : Error struct for usecase layer
type Error struct {
	Err  error // error(wrapped error)
	Kind int   // kind of error
}

// Error : get error message
func (e Error) Error() string {
	return e.Err.Error()
}
