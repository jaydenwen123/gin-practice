package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaydenwen123/go-util"
)

func main() {
	r := gin.Default()

	r.GET("/ping", pingHandler)

	//web中有三大类参数：
	//1.通过url地址后面？后指定的key、value的query 参数--->query string
	//2.post 提交的表单信息、或者json数据-->form string
	//3.uri中携带的参数，如user/1这种---> params

	r.GET("/hello", helloHandler)
	r.POST("/login", loginHandler)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func loginHandler(c *gin.Context) {
	//获取登录参数信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username:", username, "\tpassword:", password)
	if username == password {
		c.JSON(http.StatusOK, gin.H{
			"login":  "success",
			"status": "logined",
		})
		return
	}
	c.JSON(http.StatusOK, "failed")
}

func helloHandler(c *gin.Context) {
	fmt.Println("c.Params:", util.Obj2JsonStr(c.Params))
	fmt.Println("c.Keys:", util.Obj2JsonStr(c.Keys))
	fmt.Println("c.Accepted:", util.Obj2JsonStr(c.Accepted))
	c.JSON(http.StatusOK, gin.H{
		"hello":  "world",
		"author": "fly",
	})
}
