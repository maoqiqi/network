package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	// r := gin.Default()
	r := gin.New()
	// Ping test
	r.GET("/ping", ping)
	r.POST("/binding", binding)
	r.POST("/form", form)
	r.POST("/query", query)
	r.GET("/xml", xml)
	r.GET("/yaml", yaml)
	r.POST("/upload", upload)
	r.POST("/uploads", uploads)
	_ = r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

type LoginForm struct {
	User     string `json:"user" form:"user"`
	Password string `json:"password" form:"password"`
	UserName string `json:"user_name" form:"user_name"`
}

// curl -v --form user=user --form password=password http://localhost:8080/binding
func binding(c *gin.Context) {
	var form LoginForm
	// 你可以使用显式绑定声明绑定 multipart form：
	// c.ShouldBindWith(&form, binding.Form)
	// 或者简单地使用ShouldBind方法自动绑定：在这种情况下,将自动选择合适的绑定
	if c.ShouldBind(&form) == nil {
		c.JSON(200, form)
	}

	/*if err := c.BindJSON(form); err != nil {
		fmt.Println("err:", err)
		c.String(200, "fail")
	} else {
		c.JSON(200, form )
	}*/
}

// curl -v --form user=user --form password=password http://localhost:8080/form
func form(c *gin.Context) {
	user := c.PostForm("user")
	password := c.PostForm("password")
	userName := c.DefaultPostForm("user_name", "golang")

	// gin.H 是 map[string]interface{} 的一种快捷方式
	c.JSON(200, gin.H{
		"user":      user,
		"password":  password,
		"user_name": userName,
	})
}

// curl -v --form user=user --form password=password http://localhost:8080/query?id=1234&page=1
func query(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	user := c.PostForm("user")
	password := c.PostForm("password")

	c.JSON(200, gin.H{
		"id":       id,
		"page":     page,
		"user":     user,
		"password": password,
	})
}

func xml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"status": http.StatusOK, "message": "ok"})
}

func yaml(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"status": http.StatusOK, "message": "ok"})
}

// curl -X POST http://localhost:8080/upload -F "file=@/Users/mac/Downloads/csb-db-2019-10-24.txt" -H "Content-Type: multipart/form-data"
func upload(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")

	// 上传文件至指定目录
	// c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

// curl -X POST http://localhost:8080/uploads \
// -F "upload[]=@/Users/mac/Downloads/csb-db-2019-10-24.txt" \
// -F "upload[]=@/Users/mac/Downloads/csb-db-2019-10-24.txt" \
// -H "Content-Type: multipart/form-data"
func uploads(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		fmt.Println(file.Filename)

		// 上传文件至指定目录
		// c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
