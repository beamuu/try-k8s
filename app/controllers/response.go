package controllers

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Payload interface{} `json:"payload"`
}


func CreateErrorResponse(err error) Response {
	return Response{
		Success: false,
		Message: err.Error(),
		Payload: nil,
	}
}

func CreateSuccessResponse(payload interface{}) Response {
	return Response{
		Success: true,
		Message: "",
		Payload: payload,
	}
}