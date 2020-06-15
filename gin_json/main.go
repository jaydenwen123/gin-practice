package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()
	//注册处理的uri
	g.POST("/login", loginPostFormHandler)

	g.POST("/register", registerPostBodyHandler)

	g.Run(":8080")

}

//
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// curl http://localhost:8080/register -XPOST -d '{username:wenxiaofei,password:wsm6224291992,email:hello@qq.com}' -H 'Content-Type:application/json'
func registerPostBodyHandler(c *gin.Context) {
	//从post body中提取参数
	var  user   User
	c.ShouldBind(&user)
	username := c.PostForm("username")
	password := c.DefaultPostForm("password", "123456")
	email, ok := c.GetPostForm("email")
	if !ok {
		email = "unknown email"
	}
	c.JSON(http.StatusOK, gin.H{
		"action":   "register",
		"username": username,
		"password": password,
		"email":    email,
		"user":user,
	})
}

// curl http://localhost:8080/login -XPOST --form username=wenxiaofie --form password=abcdefg --form email=2282186474@qq.com
// curl http://localhost:8080/login -XPOST --form username=wenxiaofie --form password=abcdefg --form email=2282186474@qq.com
// curl http://localhost:8080/login -XPOST --form username=wenxiaofie --form password=abcdefg --form email=2282186474@qq.com
func loginPostFormHandler(c *gin.Context) {
	//从表单提取参数
	username := c.PostForm("username")
	password := c.DefaultPostForm("password", "123456")
	email, ok := c.GetPostForm("email")
	if !ok {
		email = "unknown email"
	}
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
		"email":    email,
	})
}

//./fill_sample.sh "curl http://localhost:8080/login -XPOST --form username=wenxiaofie --form password=abcdefg --form email=2282186474@qq.com"
