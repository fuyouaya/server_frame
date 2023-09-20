package core

import (
	"github.com/gin-gonic/gin"
	"server_frame/pkg/errs"
)

type HandlerFunc func(c *gin.Context) (err error)

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := h(c)
		if err != nil {
			errF := err.(*errs.Error)
			c.JSON(errF.Status, ErrRes{
				Error:     errF.Message,
				ErrorCode: errF.ErrorCode,
				Info:      errF.Info,
			})
			c.Abort()
		}
	}
}

type ErrRes struct {
	// tid TraceID 链路id
	TID string `json:"requestId"`
	// 错误响应的参考消息。前端可使用error来做提示
	Error string `json:"error"`
	// 错误码
	ErrorCode int `json:"errorCode"`
	// 用于dev调试，但这个后续看线上要用。例如一个参数错误，可能有多种原因
	Info string `json:"info,omitempty"`
}
