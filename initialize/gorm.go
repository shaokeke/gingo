package initialize

import (
	"github.com/songcser/gingo/config"
	"github.com/songcser/gingo/internal/app"
	"github.com/songcser/gingo/internal/user"
	"github.com/songcser/gingo/pkg/auth"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	switch config.GVA_CONFIG.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	// db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate
	err := db.AutoMigrate(
		// 系统模块表
		auth.BaseUser{},
		user.User{},
		app.App{}, // app表注册
	)
	if err != nil {
		os.Exit(0)
	}
}
