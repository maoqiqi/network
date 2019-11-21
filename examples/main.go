package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"network"
	"time"
)

var CheckHandler network.HandlerFunc = func(c *network.Context) network.Response {
	return c.String("ok")
}

var TestHandler network.HandlerFunc = func(c *network.Context) network.Response {
	return c.Success(nil)
}

// curl -X POST http://localhost:8080/login -H "Content-Type:application/json" -d '{"user_name":"123456","password":"123456"}'
var LoginHandler network.HandlerFunc = func(c *network.Context) network.Response {
	p := &struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}{}

	if err := c.ShouldBindJSON(p); err != nil {
		return c.Fail(201, err)
	}

	fmt.Println(p)
	// 模拟登录验证
	if p.UserName != "123456" || p.Password != "123456" {
		return c.Fail(201, "用户名或密码错误")
	}

	return c.JSON(p)
}

func Route(r *gin.Engine) {
	r.Any("/heartbeat/check", network.Handle(CheckHandler))
	r.GET("/test", network.Handle(TestHandler))
	r.POST("/login", network.Handle(LoginHandler))
}

func main() {
	//r := gin.Default()
	//Route(r)
	//_ = r.Run(":8080")

	currentTime := time.Now()
	fmt.Println(currentTime.AddDate(0, 0, 7))
	fmt.Println(currentTime.AddDate(0, 0, 7))

	a := &struct {
		t time.Time
	}{t: currentTime.AddDate(0, 0, 7)}
	fmt.Println(a.t)
}
