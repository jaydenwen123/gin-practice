package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()
	//注册处理的uri
	g.GET("/user", getUserInfo)

	g.POST("/article", postArticleInfo)

	g.Run(":8080")

}

type Article struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
}

//Bind 指定form标签时，可以将query string和form表单信息做绑定
// curl  -XPOST http://localhost:8080/article?username=123&password=23&email=2282186474@qq.com&id=123&name=test
// curl  -XPOST http://localhost:8080/article?id=123&name=23234

//BindQuery 只会绑定query参数
func postArticleInfo(c *gin.Context) {
	var article Article
	if err := c.Bind(&article); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"article": article,
		})
	} else {
		c.String(http.StatusOK, err.Error())
	}
}

//
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func getUserInfo(c *gin.Context) {
	username := c.Query("username")
	password := c.DefaultQuery("password", "123456")
	email, ok := c.GetQuery("email")
	if !ok {
		email = "unknown email"
	}
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
		"email":    email,
	})
}
