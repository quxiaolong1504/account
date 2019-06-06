package errs

import (
	"github.com/jinzhu/copier"
	"net/http"
)

type BaseError struct {
	Reason string `json:"reason"`
}

func (b BaseError) Error() string {
	return b.Reason
}

type DetailError struct {
	Code         int          `json:"code"`
	Name         string       `json:"name"`
	Message      string       `json:"message"`
	Data         *interface{} `json:"data,omitempty"`
	DebugMessage *string      `json:"debug_message,omitempty"`
}

type APIError struct {
	DetailError    DetailError `json:"error"`
	HttpStatusCode int
}

func (a APIError) Error() string {
	return ""
}

func (a *APIError) WithData(data interface{}) *APIError {
	err := &APIError{}
	copier.Copy(&err, &a)
	err.DetailError.Data = &data
	return err
}

func NewDefaultAPIError(statusCode int, errorCode int, kind, message string) *APIError {
	return &APIError{
		DetailError: DetailError{
			Code:         errorCode,
			Name:         kind,
			Message:      message,
			Data:         nil,
			DebugMessage: nil,
		},
		HttpStatusCode: statusCode,
	}
}

var (
	BadArgs = NewDefaultAPIError(http.StatusBadRequest, 4000, "BadArguments", "参数不合法")

	DigitalFailed = NewDefaultAPIError(http.StatusBadRequest, 4020, "DigitalFailed", "验证码发送失败")

	VerifyDigitalFailed = NewDefaultAPIError(http.StatusBadRequest, 4021, "VerifyDigitalFailed", "验证码不正确")

	AlreadyExist= NewDefaultAPIError(http.StatusBadRequest, 4031, "AlreadyExist", "已存在")

	WeChatLoginFailed = NewDefaultAPIError(http.StatusInternalServerError, 4030, "WeChatLoginFailed", "微信登录失败")
)
