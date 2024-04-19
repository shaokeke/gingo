package user

import (
	"github.com/songcser/gingo/pkg/model"
	"github.com/songcser/gingo/utils"
)

type User struct {
	model.BaseModel
	Name     string         `json:"name" form:"name" gorm:"column:name;type:varchar(255);not null" admin:"type:input;name:name;label:姓名"`
	Password string         `json:"password" form:"password" gorm:"column:password;type:varchar(64);not null" admin:"type:input;name:password;label:密码"`
	Email    string         `json:"email" form:"email" gorm:"column:email;type:varchar(32);not null" admin:"type:input;name:email;label:邮箱"`
	Enable   int            `json:"enable" form:"enable" gorm:"column:enable;default:1;comment:用户是否被冻结" admin:"type:radio;enum:1,2;label:禁用"`
	Geom     utils.GeoPoint `gorm:"type:geometry"`
}
