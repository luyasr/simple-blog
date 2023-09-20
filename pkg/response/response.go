package response

import "github.com/luyasr/simple-blog/pkg/e"

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func NewResponse(data any) *Response {
	return &Response{
		Code:    0,
		Data:    data,
		Message: "success",
	}
}

func NewResponseWithError(err error) *Response {
	return &Response{
		Code:    e.NewCode(err),
		Data:    nil,
		Message: e.NewMessage(err),
	}
}
