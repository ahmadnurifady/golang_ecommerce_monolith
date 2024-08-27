package dto

type BaseResponse struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_message"`
	Data            any    `json:"data"`
}
