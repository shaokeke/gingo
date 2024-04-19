package user

import (
	"github.com/gin-gonic/gin"
	"github.com/songcser/gingo/pkg/model"
	"github.com/songcser/gingo/pkg/service"
	"github.com/songcser/gingo/utils"
)

type Service struct {
	service.Service[User]
}

func NewService(u User) Service {
	return Service{service.NewBaseService[User](u)}
}

func (s Service) MakeMapper(c *gin.Context) model.Mapper[User] {
	var r Request
	err := c.ShouldBindQuery(&r)
	utils.CheckError(err)
	w := model.NewWrapper()
	w.Like("name", r.Name)
	w.Eq("Enable", 1)
	m := model.NewMapper[User](User{}, w)
	return m
}

func (s Service) MakeResponse(val model.Model) any {
	u := val.(User)
	res := Response{
		Name:   u.Name,
		Email:  u.Email,
		Enable: u.Enable,
		Geom:   u.Geom,
	}
	return res
}

func (s Service) Hello() string {
	return "Hello World"
}
