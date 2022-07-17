package server

import (
	"os"
	"tea/api"
	"tea/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}

	//仓库接口
	warehouse := r.Group("/warehouse")
	{
		warehouse.GET("ping", api.Ping)
		warehouse.POST("/list", api.WarehouseList)
	}

	//产品接口
	product := r.Group("/product")
	{
		product.GET("ping", api.Ping)
		product.POST("add", api.ProductAdd)
	}

	return r
}
