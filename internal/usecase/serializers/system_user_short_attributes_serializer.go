package serializers

import (
	"gitlab.com/abiewardani/scaffold/internal/domain"
)

type SystemUserShortAttributesSerializer struct {
	ID                string `json:"id"`
	EmailStatus       int    `json:"email_status"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	MobileCountryCode string `json:"mobile_country_code"`
	MobileNo          string `json:"mobile_no"`
	ProfilePicture    string `json:"profile_picture"`
	Email             string `json:"email"`
}

func NewSystemUserShortAttributesSerializer(data *domain.SystemUser) *SystemUserShortAttributesSerializer {
	res := SystemUserShortAttributesSerializer{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		MobileNo:  data.MobileNo(),
		Email:     data.Email(),
	}

	return &res
}
