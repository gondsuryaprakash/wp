package notification

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Model   interface{} `json:"model"`
}

func NewResponse(code, message string) Response {
	return Response{
		Code:    code,
		Message: message,
		Model:   make(map[string]interface{}),
	}
}

func NewResponseWithModel(code, message string, model interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Model:   model,
	}
}
