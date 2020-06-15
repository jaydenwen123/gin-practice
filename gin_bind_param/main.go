package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jaydenwen123/go-util"
)

func main() {
	g := gin.Default()
	//注册处理的uri
	g.Any("/user", getUserInfo)

	g.Run(":8080")

}

// curl http://localhost:8080/user?username=123&password=234 -XGET
// curl http://localhost:8080/user?username=123&password=234 -XPOST
// curl http://localhost:8080/user -XGET --data '{username:hello,password:world}' -H 'Content-Type:application/json'
// curl http://localhost:8080/user?username=123&password=123 -XGET
func getUserInfo(c *gin.Context) {
	var user User
	fmt.Println("username:",c.Query("username"))
	//ShouldBindQuery 通过form标签指定
	//BindJSON 通过json标签映射，使用Content-Type:application/json
	//Bind 绑定query string或者表单数据,form标签映射
	if err := c.Bind(&user); err != nil {
		logs.Error("ShouldBindQuery error:%s", err.Error())
		c.String(200, "error")
		return
	}
	logs.Debug(util.Obj2JsonStr(user))
	c.JSON(200, gin.H{
		"user": user,
	})
}

//
type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}

//./fill_sample.sh "curl http://localhost:8080/login -XPOST --form username=wenxiaofie --form password=abcdefg --form email=2282186474@qq.com"
