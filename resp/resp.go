package resp

type CommonResponse struct {
	Prompts string      `json:"prompts"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewStdResp(code IErrorCode, data interface{}) *CommonResponse {
	return &CommonResponse{
		Prompts: code.GetPrompts(),
		Message: "",
		Status:  0,
		Data:    data,
	}
}
