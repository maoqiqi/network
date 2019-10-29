package network

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type Response interface {
	Render()
}

type StringResponse struct {
	Context    *gin.Context   `json:"-"`
	HttpStatus int            `json:"-"`
	Message    string         `json:"message"`
	Data       [] interface{} `json:"data"`
}

func (r *StringResponse) Render() {
	r.Context.String(r.HttpStatus, r.Message, r.Data...)
}

type JSONResponse struct {
	Context    *gin.Context `json:"-"`
	HttpStatus int          `json:"-"`
	Data       interface{}  `json:"data"`
}

func (r *JSONResponse) Render() {
	r.Context.Render(r.HttpStatus, render.JSON{Data: r.Data})
}

type ApiResponse struct {
	Context    *gin.Context `json:"-"`
	HttpStatus int          `json:"-"`
	Code       int          `json:"code"`
	Msg        string       `json:"msg"`
	Data       interface{}  `json:"data"`
}

func (r *ApiResponse) Render() {
	r.Context.Render(r.HttpStatus, render.JSON{Data: r})
}

type RedirectResponse struct {
	Context    *gin.Context `json:"-"`
	HttpStatus int          `json:"-"`
	Location   string       `json:"-"`
}

func (r *RedirectResponse) Render() {
	r.Context.Redirect(r.HttpStatus, r.Location)
}

type HTMLResponse struct {
	Context    *gin.Context   `json:"-"`
	HttpStatus int            `json:"-"`
	Message    string         `json:"message"`
	Data       [] interface{} `json:"data"`
}

func (r *HTMLResponse) Render() {
	r.Context.HTML(r.HttpStatus, r.Message, r.Data)
}
