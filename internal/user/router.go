package user

import (
	"github.com/gin-gonic/gin"
	"github.com/songcser/gingo/pkg/router"
)

func InitRouter(g *gin.RouterGroup) {
	// 组合router
	r := router.NewRouter(g.Group("user"))
	// 组合api curd 和 BaseService
	a := NewApi()
	// 绑定api
	r.BindApi("", a)
	// 自定义api
	r.BindGet("hello", a.Hello) //
}
