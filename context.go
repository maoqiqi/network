package network

import (
	"github.com/gin-gonic/gin"
)

var SuccessCode int = 10000

type Context struct {
	*gin.Context
	HttpStatus int
	Response   Response
}

func (c *Context) GetToken() string {
	return c.Request.Header.Get("Access-Token")
}

func (c *Context) Fail(code int, msg interface{}) Response {
	var message string
	if m, ok := msg.(error); ok {
		message = m.Error()
	} else {
		message = ToStr(msg)
	}
	return &ApiResponse{
		Context:    c.Context,
		HttpStatus: getHttpStatus(c, 200),
		Code:       code,
		Msg:        message,
		Data:       "",
	}
}

func (c *Context) Success(data interface{}) Response {
	return &ApiResponse{
		Context:    c.Context,
		HttpStatus: getHttpStatus(c, 200),
		Code:       SuccessCode,
		Msg:        "success",
		Data:       IfNull(data),
	}
}

func (c *Context) String(format string, values ...interface{}) Response {
	return &StringResponse{
		Context:    c.Context,
		HttpStatus: getHttpStatus(c, 200),
		Message:    format,
		Data:       values,
	}
}

func (c *Context) JSON(data interface{}) Response {
	return &JSONResponse{
		Context:    c.Context,
		HttpStatus: getHttpStatus(c, 200),
		Data:       IfNull(data),
	}
}

func (c *Context) Redirect(location string) Response {
	return &RedirectResponse{
		Context:    c.Context,
		HttpStatus: getHttpStatus(c, 302),
		Location:   location,
	}
}

func getHttpStatus(c *Context, status int) int {
	if c.HttpStatus == 0 {
		return status
	}
	return c.HttpStatus
}

func IfNull(data interface{}) interface{} {
	if data == nil {
		return ""
	}
	return data
}
