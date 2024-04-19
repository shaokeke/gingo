package user

import (
	"github.com/songcser/gingo/pkg/admin"
)

func Admin() {
	var u User
	admin.New(u, "user", "用户")
}
