package utils

import (
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/abiewardani/scaffold/pkg/shared"
)

// JSONSuccess ...
func JSONSuccess(e echo.Context, params interface{}) error {
	return e.JSON(http.StatusOK, params)
}

func JSONError(e echo.Context, err error) error {
	if shared.IsMultiStringValidationError(err) {
		return e.JSON(http.StatusUnprocessableEntity, err.(*shared.MultiStringValidationError))
	}

	if shared.IsMultiStringBadRequestError(err) {
		return e.JSON(http.StatusBadRequest, err.(*shared.MultiStringBadRequestError))
	}

	if shared.IsNewMultiStringForbidenErrorWithData(err) {
		return e.JSON(http.StatusForbidden, err.(*shared.MultiStringForbidenErrorWithData))
	}

	if shared.IsMultiStringUnauthorizedError(err) {
		return e.JSON(http.StatusUnauthorized, shared.DefaultMultiStringUnauthorizedError())
	}

	if shared.IsMultiStringHTTPErrorDataNotFound(err) {
		return e.JSON(shared.HTTPErrorDataNotFound, shared.DefaultMultiStringHTTPErrorDataNotFound())
	}

	return e.JSON(http.StatusBadRequest, shared.DefaultMultiStringInternalServerError())
}
