package network

import "github.com/gin-gonic/gin"

type IHandler interface {
	Handle(c *Context) Response
}

type HandlerFunc func(c *Context) Response

func (h HandlerFunc) Handle(c *Context) Response {
	return h(c)
}

func Handle(handler IHandler) gin.HandlerFunc {
	return func(context *gin.Context) {
		c := getContext(context)
		if resp := handler.Handle(c); resp != nil {
			c.Abort()
			c.Response = resp
			c.Response.Render()
		}
	}
}

func getContext(c *gin.Context) *Context {
	return &Context{Context: c}
}
