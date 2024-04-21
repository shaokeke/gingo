package user

import (
	"github.com/gin-gonic/gin"
	"github.com/songcser/gingo/pkg/api"
	"github.com/songcser/gingo/pkg/response"
)

type Api struct {
	api.Api
	Service Service
}

func NewApi() Api {
	var user User
	s := NewService(user)
	baseApi := api.NewApi[User](s)
	return Api{Api: baseApi, Service: s}
}

// Hello
// @Tags 用户管理
// @Summary TestHello
// @Success 200 {string} welcome
// @Router /user/hello [get]
func (a Api) Hello(c *gin.Context) {
	str := a.Service.Hello()
	response.OkWithData(str, c)
}
