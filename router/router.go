package router

import (
	"xfire/global"
	"xfire/server/factory"
	"xfire/server/news"

	"github.com/gin-gonic/gin"
)

func FactoryRouter(r *gin.Engine) {
	api := new(factory.FactoryApi)
	factoryRouter := r.Group("factory")
	factoryRouter.POST("createBrand", api.CreateBrand)
	factoryRouter.POST("createBuild", api.CreateBuild)
	factoryRouter.POST("createGoods", api.CreateGoods)

	factoryRouter.POST("deleteBrandByIds", api.DeleteBrandByIds)
	factoryRouter.POST("deleteBuildByIds", api.DeleteBuildByIds)
	factoryRouter.POST("deleteGoodsByIds", api.DeleteGoodsByIds)

	factoryRouter.POST("getBrandList", api.GetBrandList)
	factoryRouter.POST("getBuildList", api.GetBuildList)
	factoryRouter.POST("getGoodsList", api.GetGoodsList)

	factoryRouter.POST("updateBrand", api.UpdateBrand)
	factoryRouter.POST("updateBuild", api.UpdateBuild)
	factoryRouter.POST("updateGoods", api.UpdateGoods)
	factoryRouter.POST("showCost", api.ShowCost)
}

func NewRouter(r *gin.Engine) {
	api := new(news.News)
	newsRouter := r.Group("news")
	newsRouter.POST("getGoodsList", api.GetGoodsList)
}

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
	FactoryRouter(Router)
	NewRouter(Router)
	global.LOG.Info("router register success")
	return Router
}
