package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/abiewardani/scaffold/pkg/shared"
)

type SystemUserLoginParams struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

func (c *SystemUserLoginParams) Validate() error {
	if err := validation.Validate(c.UsernameOrEmail, validation.Required); err != nil {
		return shared.NewMultiStringValidationError(shared.HTTPErrorUnprocessableEntity, `Username Or Email Is Required`)
	}

	if err := validation.Validate(c.Password, validation.Required); err != nil {
		return shared.NewMultiStringValidationError(shared.HTTPErrorUnprocessableEntity, `Password Is Required`)
	}

	return nil
}
