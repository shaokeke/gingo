package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/songcser/gingo/config"
	"github.com/songcser/gingo/internal/app"
	"github.com/songcser/gingo/internal/user"
	"github.com/songcser/gingo/middleware"
	"github.com/songcser/gingo/utils"
)

func Routers() *gin.Engine {

	if err := utils.Translator("zh"); err != nil {
		config.GVA_LOG.Error(err.Error())
		return nil
	}
	// 在初始化 gin 路由器之前，您必须调用 SetMode 方法
	if !config.GVA_CONFIG.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	//gin.SetMode(gin.DebugMode)
	Router := gin.Default()

	Swagger(Router)
	Router.Use(middleware.Recovery())
	Router.Use(middleware.Logger())
	HealthGroup := Router.Group("")
	{
		// 健康监测
		HealthGroup.GET("/health", HealthCheck)
	}

	ApiGroup := Router.Group(config.GVA_CONFIG.RouterPrefix)
	app.InitRouter(ApiGroup)
	user.InitRouter(ApiGroup)
	return Router
}

// @Summary 健康检查
// @Tags 健康检查T
// @Success 200 {string} welcome
// @Router /health [get]
func HealthCheck(g *gin.Context) {
	g.JSON(http.StatusOK, "ok...")
}
