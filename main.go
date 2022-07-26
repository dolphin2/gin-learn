package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/**
首先，我们使用了gin.Default()生成了一个实例，这个实例即 WSGI 应用程序。
接下来，我们使用r.Get("/", ...)声明了一个路由，告诉 Gin 什么样的URL 能触发传入的函数，这个函数返回我们想要显示在用户浏览器中的信息。
最后用 r.Run()函数来让应用运行在本地服务器上，默认监听端口是 _8080_，可以传入参数设置端口，例如r.Run(":9999")即运行在 _9999_端口。
*/
func main() {
	r := gin.Default()
	getPostTest(r)
	groupRoutes(r)
	r.Run()
}

func getPostTest(r *gin.Engine) {
	//r.GET("/", func(c *gin.Context) {
	//	c.String(200, "hello doudou")
	//})
	//匹配/user/doudou
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s", name)
	})
	//获取Query参数
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "nurse")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})
	//post 请求
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	//Query和POST混合参数
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000")
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})
	//重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})
	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})
}

func uploadFile(r *gin.Engine) {
	r.POST("upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})

}

func uploadFiles(r *gin.Engine) {
	// Multipart form
	r.POST("upload2", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
		}
		c.String(http.StatusOK, "%d files uploaded!", len(files))
	})

}
