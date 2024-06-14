package helpers

type CommonResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GenerateErrorWithMessage(message string) CommonResponse {
	return CommonResponse{
		message, nil,
	}
}

func GenerateSuccessWithData(message string, data interface{}) CommonResponse {
	return CommonResponse{
		message, data,
	}
}
