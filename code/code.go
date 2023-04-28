package code

// ErrorCode is an abstraction of the type of error that occurred.
// By comparing ErrorCode, it's possible to determine at which
// layer the error occurred, which is useful for generationg
// error messages for responses.
type ErrorCode string

var (
	CodeOK ErrorCode = "ok"

	CodeUnknown ErrorCode = "unknown error"
)

// Error returns error message.
// It's an implementation of the error interface.
// Use when comparing code with errors.Is().
func (c ErrorCode) Error() string {
	return string(c)
}

var _ error = (*ErrorCode)(nil)
