package apperror

import "net/http"

var (
	ErrInvalidRequest     = &apiError{statusCode: http.StatusBadRequest, CodeText: "invalid request"}
	ErrUnauthorizedApiKey = &apiError{statusCode: http.StatusUnauthorized, CodeText: "The API key provided is not currently able to query this endpoint"}
	ErrInvalidToken       = &apiError{statusCode: http.StatusUnauthorized, CodeText: "Invalid API key or authorization header"}
	ErrNotFound           = &apiError{statusCode: http.StatusNotFound, CodeText: "Resource Not Found"}
	ErrInternalError      = &apiError{statusCode: http.StatusInternalServerError, CodeText: "Something went wrong internally, please try again later"}
)

type APIError interface {
	Message() (int, string)
}

type apiError struct {
	statusCode int
	CodeText   string
}

func (e *apiError) Message() (int, string) {
	return e.statusCode, e.CodeText
}

type apiErrorWrapError struct {
	apiErr *apiError
	err    error
}

func (e *apiErrorWrapError) Error() string {
	return e.err.Error()
}

func (e *apiErrorWrapError) Message() (int, string) {
	return e.apiErr.Message()
}

func WrapError(err error, apiErr *apiError) error {
	return &apiErrorWrapError{err: err, apiErr: apiErr}
}
