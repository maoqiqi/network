package main

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/url"
	"network"
)

var TestHandler network.HandlerFunc = func(c *network.Context) network.Response {
	_ = c.Request.ParseForm()
	_ = c.Request.ParseMultipartForm(32 << 20)

	// 地址问号后的参数
	queryForm, _ := url.ParseQuery(c.Request.URL.RawQuery)
	// 存储了get,post,put参数,在使用之前需要调用ParseForm方法
	form := c.Request.Form
	// 存储了post,put参数,在使用之前需要调用ParseForm方法
	postForm := c.Request.PostForm
	// 存储了包含了文件上传的表单的post参数,在使用之前需要调用ParseMultipartForm方法
	multipartForm := c.Request.MultipartForm

	return c.JSON(&struct {
		QueryForm     map[string][]string `json:"query_form"`
		Form          url.Values          `json:"form"`
		PostForm      url.Values          `json:"post_form"`
		MultipartForm *multipart.Form     `json:"multipart_form"`
	}{
		QueryForm:     queryForm,
		Form:          form,
		PostForm:      postForm,
		MultipartForm: multipartForm,
	})
}

func main() {
	r := gin.Default()
	r.Any("/test", network.Handle(TestHandler))
	_ = r.Run(":8080")
}
