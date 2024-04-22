package server

import (
	"fisco-drug-trace/api"
	"fisco-drug-trace/middleware"
	"os"

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
		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 用户列表

		// 需要登录保护的
		common := v1.Group("")
		common.Use(middleware.AuthRequired())
		{
			// User Routing
			common.GET("user/me", api.UserMe)
			common.GET("user/logout", api.UserLogout)
			common.POST("user/update/me", api.UserUpdateMe)
			common.POST("drug/list", api.DrugList)
		}

		admin := v1.Group("")
		admin.Use(middleware.AdminAuthRequired())
		{
			admin.POST("users", api.UserList)
			admin.POST("user/update", api.UserUpdate)
		}

		merchant := v1.Group("")
		merchant.Use(middleware.MerchantAuthRequired())
		{
			merchant.POST("drug/add", api.ProduceDrug)
			merchant.POST("drug/sale", api.SellDrug)
			merchant.POST("user/dealer", api.DealerList)
		}

		dealer := v1.Group("")
		dealer.Use(middleware.DealerAuthRequired())
		{
			dealer.POST("drug/accept", api.DealerAcceptDrug)
		}

		user := v1.Group("")
		user.Use(middleware.UserAuthRequired())
		{
			user.POST("drug/buy", api.BuyDrug)
			user.POST("drug/buy/list", api.DrugPayableList)
		}

	}
	return r
}
