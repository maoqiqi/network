package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"network"
)

// 测试网络请求参数获取情况
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

	b, _ := ioutil.ReadAll(c.Request.Body)

	// 测试Cookie
	// c.SetCookie("user_id", "10001", int(60*time.Second), "/", "", false, false)

	// 测试重定向
	// return c.Redirect("/redirect")

	return c.JSON(&struct {
		QueryForm     map[string][]string `json:"query_form"`
		Form          url.Values          `json:"form"`
		PostForm      url.Values          `json:"post_form"`
		MultipartForm *multipart.Form     `json:"multipart_form"`
		Body          string              `json:"body"`
	}{
		QueryForm:     queryForm,
		Form:          form,
		PostForm:      postForm,
		MultipartForm: multipartForm,
		Body:          string(b),
	})
}

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Page     int    `json:"page"`
}

// 测试GIN框架从请求中获取参数的方法
// curl -d 'user_name=admin&password=123456' http://127.0.0.1:8080/params?user_name=zhangsan&page=1
var ParamsHandler network.HandlerFunc = func(c *network.Context) network.Response {
	// 查询请求URL后面的参数
	userName1 := c.Query("user_name")
	// 查询请求URL后面的参数,如果没有填写默认值
	page := c.DefaultQuery("page", "0")
	// 从表单中查询参数
	userName2 := c.PostForm("user_name")

	// POST和PUT主体参数优先于URL查询字符串值
	userName3 := c.Request.FormValue("user_name")

	// 返回POST并放置body参数,URL查询参数被忽略
	userName4 := c.Request.PostFormValue("user_name")

	// 从表单中查询参,如果没有填写默认值
	password := c.DefaultPostForm("password", "111111")

	var u1 User
	// 从http.Request中读取值到User结构体中,手动确定绑定类型binding.Form
	_ = c.BindWith(&u1, binding.Form)

	var u2 User
	// 从http.Request中读取值到User结构体中,根据请求方法类型和请求内容格式类型自动确定绑定类型
	_ = c.Bind(&u2)

	fmt.Println("userName1:", userName1)
	fmt.Println("page:", page)
	fmt.Println("userName2:", userName2)
	fmt.Println("userName3:", userName3)
	fmt.Println("userName4:", userName4)
	fmt.Println("password:", password)
	fmt.Printf("u1:%v\n", u1)
	fmt.Printf("u2:%v\n", u2)

	/*
		userName1: zhangsan
		page: 1
		userName2: admin
		userName3: admin
		userName4: admin
		password: 123456
		u1:{  0}
		u2:{  0}
	*/

	return c.JSON(nil)
}

func main() {
	r := gin.Default()
	r.Any("/test", network.Handle(TestHandler))
	r.Any("/redirect", func(c *gin.Context) {
		c.String(http.StatusOK, "redirect")
	})
	r.Any("/params", network.Handle(ParamsHandler))
	_ = r.Run(":8080")
}
