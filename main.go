package main

import (
	"fmt"
	"os"
	"time"
	"xfire/global"
	"xfire/router"

	"xfire/docs"

	swaggerFiles "github.com/swaggo/files"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {

	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

// @title                       进销存 Swagger API接口文档
// @version                     v0.0.1
// @description                 gin+vue
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	docs.SwaggerInfo.Title = "drp"
	global.InitConfig(os.Args[1])
	global.InitLogger()
	global.InitDb(global.C.Dblink)
	Router := router.Routers()

	// use ginSwagger middleware to serve the API docs
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	address := fmt.Sprintf(":%d", global.C.Addr)
	s := initServer(address, Router)
	global.LOG.Info(fmt.Sprintf("swagger addr: %s/swagger/index.html", address))

	global.LOG.Info("server run success on ", zap.String("address", address))

	global.LOG.Error(s.ListenAndServe().Error())
	// cmd.RunPrompt()
}
