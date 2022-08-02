package shared

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	// HTTPErrorBadRequest bad request code
	HTTPErrorBadRequest = http.StatusBadRequest
	// HTTPErrorNotFound not found code
	HTTPErrorNotFound = http.StatusNotFound
	// HTTPErrorUnauthorized unauthorized code
	HTTPErrorUnauthorized = http.StatusUnauthorized
	// HTTPErrorForbidden forbidden code
	HTTPErrorForbidden = http.StatusForbidden
	// HTTPErrorMethodNotAllowed method not allowed code
	HTTPErrorMethodNotAllowed = http.StatusMethodNotAllowed
	// HTTPErrorInternalServer internal error code
	HTTPErrorInternalServer = http.StatusInternalServerError
	// HTTPErrorUnprocessableEntity unprocessable entity code
	HTTPErrorUnprocessableEntity = http.StatusUnprocessableEntity
	// HTTPErrorTimeOut response timeout
	HTTPErrorTimeOut = http.StatusRequestTimeout // use this error if net/http request canceled/response timeout
	// HTTPErrorDataNotFound Application/data specific
	HTTPErrorDataNotFound = 100 // use this error if requested data not found
)

var messageMap = map[int]string{
	http.StatusBadRequest:          "Bad request",
	http.StatusNotFound:            "Route not found",
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusForbidden:           "Forbidden",
	http.StatusMethodNotAllowed:    "Method not allowed",
	http.StatusInternalServerError: "Internal server error",
	http.StatusUnprocessableEntity: "Invalid inputs",
	HTTPErrorDataNotFound:          "Data not found",
	http.StatusRequestTimeout:      "Request time out",
}

// StringMap get message text from constant code and return string map contain multilang string.
// It returns empty if the code unknown.
func StringMap(code int) string {
	return messageMap[code]
}

type MultiLangError struct {
	Errors    string    `json:"errors"`
	ErrorCode ErrorCode `json:"error_code"`
}

type ErrorCode struct {
	ErrorID     int         `json:"id"`
	Data        interface{} `json:"data,omitempty"`
	Description string      `json:"description"`
}

type MultiLangErrorWithData struct {
	ErrorID int               `json:"code"`
	Msg     map[string]string `json:"message"`
	Data    interface{}       `json:"data"`
}

// MultiStringValidationError multi string error will contain multi lang error message.
type MultiStringValidationError MultiLangError

// NewMultiStringValidationError create new multi string error struct
func NewMultiStringValidationError(code int, msg string) *MultiStringValidationError {
	return &MultiStringValidationError{
		Errors: msg,
		ErrorCode: ErrorCode{
			ErrorID:     code,
			Description: msg,
		},
	}
}

func (c *MultiStringValidationError) Code() int {
	return c.ErrorCode.ErrorID
}

func (c *MultiStringValidationError) Message() string {
	return c.Errors
}

// Error comply with error interface
func (c *MultiStringValidationError) Error() string {
	// nolint
	b, _ := json.Marshal(c)

	return string(b)
}

// IsMultiStringValidationError check whether error is MultiStringValidationError pointer
func IsMultiStringValidationError(err error) bool {
	switch err.(type) {
	case *MultiStringValidationError:
		return true
	default:
		return false
	}
}

// MultiStringBadRequestError multi string bad request error will contain multi lang error message.
type MultiStringBadRequestError MultiLangError

// NewMultiStringBadRequestError create new multi string bad request error
func NewMultiStringBadRequestError(code int, msg string) *MultiStringBadRequestError {
	return &MultiStringBadRequestError{
		Errors: msg,
		ErrorCode: ErrorCode{
			ErrorID:     code,
			Description: msg,
		},
	}
}

// MultiStringBadRequestError multi string bad request error will contain multi lang error message.
type MultiStringForbidenErrorWithData MultiLangErrorWithData

// NewMultiStringBadRequestError create new multi string bad request error
func NewMultiStringForbidenErrorWithData(code int, msg map[string]string, data interface{}) *MultiStringForbidenErrorWithData {
	return &MultiStringForbidenErrorWithData{
		ErrorID: code,
		Msg:     msg,
		Data:    data,
	}
}

func (c *MultiStringForbidenErrorWithData) Code() int {
	return c.ErrorID
}

func (c *MultiStringForbidenErrorWithData) Result() interface{} {
	return c.Data
}

func (c *MultiStringForbidenErrorWithData) Message() map[string]string {
	return c.Msg
}

// Error comply with error interface
func (c *MultiStringForbidenErrorWithData) Error() string {
	// nolint
	b, _ := json.Marshal(c)

	return string(b)
}

// DefaultMultiStringBadRequestError ...
func DefaultMultiStringBadRequestError() *MultiStringBadRequestError {
	return &MultiStringBadRequestError{
		Errors: StringMap(http.StatusBadRequest),
		ErrorCode: ErrorCode{
			ErrorID:     http.StatusBadRequest,
			Description: StringMap(http.StatusBadRequest),
		},
	}
}

// Error comply with error interface
func (c *MultiStringBadRequestError) Error() string {
	// nolint
	b, _ := json.Marshal(c)

	return string(b)
}

// IsMultiStringBadRequestError check if it was bad request error
func IsMultiStringBadRequestError(err error) bool {
	switch err.(type) {
	case *MultiStringBadRequestError:
		return true
	default:
		return false
	}
}

// IsNewMultiStringForbidenErrorWithData check if it was bad request error
func IsNewMultiStringForbidenErrorWithData(err error) bool {
	switch err.(type) {
	case *MultiStringForbidenErrorWithData:
		return true
	default:
		return false
	}
}

// MultiStringUnauthorizedError multi string unauthorized request error will contain multi lang error message.
type MultiStringUnauthorizedError MultiLangError

// NewMultiStringUnauthorizedError create new multi string bad request error
func NewMultiStringUnauthorizedError(code int, msg string) *MultiStringUnauthorizedError {
	return &MultiStringUnauthorizedError{
		Errors: msg,
		ErrorCode: ErrorCode{
			ErrorID:     code,
			Description: msg,
		},
	}
}

func DefaultMultiStringUnauthorizedError() *MultiStringUnauthorizedError {
	return &MultiStringUnauthorizedError{
		Errors: StringMap(HTTPErrorUnauthorized),
		ErrorCode: ErrorCode{
			ErrorID:     http.StatusUnauthorized,
			Description: StringMap(HTTPErrorUnauthorized),
		},
	}
}

// Error comply with error interface
func (c *MultiStringUnauthorizedError) Error() string {
	// nolint
	b, _ := json.Marshal(c)

	return string(b)
}

// IsMultiStringUnauthorizedError check if it was bad request error
func IsMultiStringUnauthorizedError(err error) bool {
	switch err.(type) {
	case *MultiStringUnauthorizedError:
		return true
	default:
		return false
	}
}

// MultiStringForbiddenError multi string forbidden request error will contain multi lang error message.
type MultiStringForbiddenError MultiLangError

// NewMultiStringForbiddenError create new multi string forbidden request error
func NewMultiStringForbiddenError(code int, msg string) *MultiStringForbiddenError {
	return &MultiStringForbiddenError{
		Errors: msg,
		ErrorCode: ErrorCode{
			ErrorID:     code,
			Description: msg,
		},
	}
}

func DefaultMultiStringForbiddenError() *MultiStringForbiddenError {
	return &MultiStringForbiddenError{
		Errors: StringMap(http.StatusForbidden),
		ErrorCode: ErrorCode{
			ErrorID:     http.StatusForbidden,
			Description: StringMap(http.StatusForbidden),
		},
	}
}

// Error comply with error interface
func (c *MultiStringForbiddenError) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// IsMultiStringForbiddenError check if it was forbidden request error
func IsMultiStringForbiddenError(err error) bool {
	switch err.(type) {
	case *MultiStringForbiddenError:
		return true
	default:
		return false
	}
}

type MultiStringInternalServerError MultiLangError

func DefaultMultiStringInternalServerError() *MultiStringInternalServerError {
	return &MultiStringInternalServerError{
		Errors: StringMap(http.StatusInternalServerError),
		ErrorCode: ErrorCode{
			ErrorID:     http.StatusInternalServerError,
			Description: StringMap(http.StatusInternalServerError),
		},
	}
}

// NewMultiStringForbiddenError create new multi string forbidden request error
func NewMultiStringInternalServerError(msg string) *MultiStringInternalServerError {
	return &MultiStringInternalServerError{
		Errors: msg,
		ErrorCode: ErrorCode{
			ErrorID:     http.StatusInternalServerError,
			Description: msg,
		},
	}
}

// Error comply with error interface
func (c *MultiStringInternalServerError) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type MultiStringRouteNotFoundError MultiLangError

func DefaultMultiStringRouteNotFoundError() *MultiStringRouteNotFoundError {
	return &MultiStringRouteNotFoundError{
		Errors: StringMap(http.StatusNotFound),
		ErrorCode: ErrorCode{
			ErrorID:     http.StatusNotFound,
			Description: StringMap(http.StatusNotFound),
		},
	}
}

type MultiStringHTTPErrorDataNotFound MultiLangError

func DefaultMultiStringHTTPErrorDataNotFound() *MultiStringHTTPErrorDataNotFound {
	return &MultiStringHTTPErrorDataNotFound{
		Errors: StringMap(HTTPErrorDataNotFound),
		ErrorCode: ErrorCode{
			ErrorID:     HTTPErrorDataNotFound,
			Description: StringMap(HTTPErrorDataNotFound),
		},
	}
}

// IsMultiStringHTTPErrorDataNotFound check if it was forbidden request error
func IsMultiStringHTTPErrorDataNotFound(err error) bool {
	switch err.(type) {
	case *MultiStringHTTPErrorDataNotFound:
		return true
	default:
		return false
	}
}

// Error comply with error interface
func (c *MultiStringHTTPErrorDataNotFound) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type MultiStringMethodNotAllowedError MultiLangError

func DefaultMultiStringMethodNotAllowedError() *MultiStringMethodNotAllowedError {
	return &MultiStringMethodNotAllowedError{
		Errors: StringMap(http.StatusMethodNotAllowed),
		ErrorCode: ErrorCode{
			ErrorID:     http.StatusMethodNotAllowed,
			Description: StringMap(http.StatusMethodNotAllowed),
		},
	}
}

type MultiStringHTTPErrorTimeOut MultiLangError

func DefaultMultiStringHTTPErrorTimeOut() *MultiStringHTTPErrorTimeOut {
	return &MultiStringHTTPErrorTimeOut{
		Errors: StringMap(http.StatusRequestTimeout),
		ErrorCode: ErrorCode{
			ErrorID:     http.StatusRequestTimeout,
			Description: StringMap(http.StatusRequestTimeout),
		},
	}
}

// IsMultiStringHTTPErrorTimeOut check if it was timeout request error
func IsMultiStringHTTPErrorTimeOut(err error) bool {
	if strings.ContainsAny(err.Error(), "timeout") {
		return true
	}
	return false
}

// Error comply with error interface
func (c *MultiStringHTTPErrorTimeOut) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type MultiStringIdError MultiLangError

func DefaultMultiStringIdError() *MultiStringIdError {
	return &MultiStringIdError{
		Errors: StringMap(http.StatusUnprocessableEntity),
		ErrorCode: ErrorCode{
			ErrorID:     http.StatusUnprocessableEntity,
			Description: StringMap(http.StatusUnprocessableEntity),
		},
	}
}

// IsMultiStringIdError check if it was timeout request error
func IsMultiStringIdError(err error) bool {
	switch err.(type) {
	case *MultiStringIdError:
		return true
	default:
		return false
	}
}

// Error comply with error interface
func (c *MultiStringIdError) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}
