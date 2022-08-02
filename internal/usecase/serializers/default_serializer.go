package serializers

type DefaultSerializer struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewDefaultSerializer(success bool, message string) *DefaultSerializer {
	res := DefaultSerializer{
		Success: success,
		Message: message,
	}

	return &res
}

func NewSuccessDefaultSerializer(message string) *DefaultSerializer {
	return NewDefaultSerializer(true, message)
}
