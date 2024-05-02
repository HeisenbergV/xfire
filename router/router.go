package router

import (
	"xfire/global"
	"xfire/server/factory"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

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
	Router.POST("createBrand", api.CreateBrand)
	Router.POST("createBuild", api.CreateBuild)
	Router.POST("createProduct", api.CreateProduct)

	Router.POST("deleteBrandByIds", api.DeleteBrandByIds)
	Router.POST("deleteBuildByIds", api.DeleteBuildByIds)
	Router.POST("deleteProductByIds", api.DeleteProductByIds)

	Router.POST("getBrandList", api.GetBrandList)
	Router.POST("getBuildList", api.GetBuildList)
	Router.POST("getProductList", api.GetProductList)

	Router.POST("updateBrand", api.UpdateBrand)
	Router.POST("updateBuild", api.UpdateBuild)
	Router.POST("updateProduct", api.UpdateProduct)
	Router.POST("showCost", api.ShowCost)

	// apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds)

	global.LOG.Info("router register success")
	return Router
}