package apperror

import "net/http"

var (
	ErrInvalidRequest = &apiError{statusCode: http.StatusBadRequest, code: "INVALID_REQUEST", description: "Invalid request"}
	// ErrUnauthorizedApiKey     = &apiError{statusCode: http.StatusUnauthorized, description: "The API key provided is not currently able to query this endpoint"}
	ErrUnauthorized          = &apiError{statusCode: http.StatusUnauthorized, code: "UNAUTHORIZED_API_KEY", description: "Invalid API key or authorization header"}
	ErrAuthorize             = &apiError{statusCode: http.StatusForbidden, code: "AUTHORIZATION_ERROR", description: "Authorization error"}
	ErrNotFound              = &apiError{statusCode: http.StatusNotFound, code: "NOT_FOUND", description: "Resource Not Found"}
	ErrInternalError         = &apiError{statusCode: http.StatusInternalServerError, code: "INTERNAL_ERROR", description: "Something went wrong internally, please try again later"}
	ErrMissingRequiredParams = &apiError{statusCode: http.StatusBadRequest, code: "MISSING_REQUIRED_PARAMS", description: "Missing required params"}
)

type APIError interface {
	Message() (statusCode int, code, description string)
}

type apiError struct {
	statusCode  int
	code        string
	description string
}

func (e *apiError) Message() (int, string, string) {
	return e.statusCode, e.code, e.description
}

func (e *apiError) Error() string {
	return e.description
}

type apiErrorWrapError struct {
	apiErr *apiError
	err    error
}

func (e *apiErrorWrapError) Error() string {
	return e.err.Error()
}

func (e *apiErrorWrapError) Message() (int, string, string) {
	return e.apiErr.Message()
}

func WrapError(err error, apiErr *apiError) error {
	return &apiErrorWrapError{err: err, apiErr: apiErr}
}
