package serializers

type LoginSerializer struct {
	User  *SystemUserShortAttributesSerializer `json:"user"`
	Token struct {
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	} `json:"tokens"`
}
