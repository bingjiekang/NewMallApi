package mallservice

import (
	_ "net/http"

	"github.com/gin-gonic/gin"
)

type MallUser struct {
}

// 注册
func (*MallUser) Register(c *gin.Context) {

}

// 登陆
func (*MallUser) Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(200, "login.html", nil)
	}
}

// 首页
func (*MallUser) Index(c *gin.Context) {
	// if
	// if c.Request.Method == "POST" {

	// } else if c.Request.Method == "GET" {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"title": "Users",
	// 	})
	// }

}
