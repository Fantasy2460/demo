package api_helper

import "errors"

// 响应结构体
type Response struct {
	Message string `json:"message"`
}

// 响应错误结构体
type ErrorResponse struct {
	Message string `json:"errorMessage"`
	Type    string `json:"type"`
}

// 自定义错误
var (
	ErrInvalidBody  = errors.New("请检查你的请求体")
	ErrInvalidToken = errors.New("Token验证失败")
)
