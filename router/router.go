package router

import (
	"xfire/global"
	"xfire/server/factory"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	// 设置为发布模式
	if global.C.Env == "public" {
		gin.SetMode(gin.ReleaseMode) //DebugMode ReleaseMode TestMode
	}

	Router := gin.New()
	Router.Use(gin.Recovery())
	if global.C.Env != "public" {
		Router.Use(gin.Logger())
	}

	api := new(factory.FactoryApi)
	Router.POST("factory/createBrand", api.CreateBrand)
	Router.POST("factory/createBuild", api.CreateBuild)
	Router.POST("factory/createGoods", api.CreateGoods)

	Router.POST("factory/deleteBrandByIds", api.DeleteBrandByIds)
	Router.POST("factory/deleteBuildByIds", api.DeleteBuildByIds)
	Router.POST("factory/deleteGoodsByIds", api.DeleteGoodsByIds)

	Router.POST("factory/getBrandList", api.GetBrandList)
	Router.POST("factory/getBuildList", api.GetBuildList)
	Router.POST("factory/getGoodsList", api.GetGoodsList)

	Router.POST("factory/updateBrand", api.UpdateBrand)
	Router.POST("factory/updateBuild", api.UpdateBuild)
	Router.POST("factory/updateGoods", api.UpdateGoods)
	Router.POST("factory/showCost", api.ShowCost)

	// apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds)

	global.LOG.Info("router register success")
	return Router
}
