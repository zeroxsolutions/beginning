package domain

type ErrorCode string

const (
	ErrorCodeInternalServerError ErrorCode = "internal_server_error"
)

type Error struct {
	Code        ErrorCode `json:"code"`
	Description string    `json:"description"`
}

type HttpError struct {
	Error
	StatusCode int `json:"status_code"`
}

func NewError(code ErrorCode, description string) *Error {
	return &Error{Code: code, Description: description}
}

func NewHttpError(code ErrorCode, description string, statusCode int) *HttpError {
	return &HttpError{Error: Error{Code: code, Description: description}, StatusCode: statusCode}
}
