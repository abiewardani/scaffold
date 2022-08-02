package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/abiewardani/scaffold/pkg/shared"
)

type SystemUserCreateParams struct {
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	MobileCountryCode string `json:"mobile_country_code"`
	MobileNo          string `json:"mobile_no"`
}

func (c *SystemUserCreateParams) Validate() error {
	if err := validation.Validate(c.Username, validation.Required); err != nil {
		return shared.NewMultiStringValidationError(shared.HTTPErrorUnprocessableEntity, `Username Is Required`)
	}

	if err := validation.Validate(c.Email, validation.Required); err != nil {
		return shared.NewMultiStringValidationError(shared.HTTPErrorUnprocessableEntity, `Email Is Required`)
	}

	if err := validation.Validate(c.Password, validation.Required); err != nil {
		return shared.NewMultiStringValidationError(shared.HTTPErrorUnprocessableEntity, `Password Is Required`)
	}

	if err := validation.Validate(c.FirstName, validation.Required); err != nil {
		return shared.NewMultiStringValidationError(shared.HTTPErrorUnprocessableEntity, `First Name Is Required`)
	}

	if err := validation.Validate(c.LastName, validation.Required); err != nil {
		return shared.NewMultiStringValidationError(shared.HTTPErrorUnprocessableEntity, `Last Name Is Required`)
	}

	if err := validation.Validate(c.MobileCountryCode, validation.Required); err != nil {
		return shared.NewMultiStringValidationError(shared.HTTPErrorUnprocessableEntity, `Mobile Country Code Is Required`)
	}

	if err := validation.Validate(c.MobileNo, validation.Required); err != nil {
		return shared.NewMultiStringValidationError(shared.HTTPErrorUnprocessableEntity, `Mobile No Code Is Required`)
	}

	return nil
}
