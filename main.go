package main

import (
	"fmt"
	"os"
	"time"
	"xfire/global"
	"xfire/router"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
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

func main() {
	global.InitConfig(os.Args[1])
	global.InitLogger()
	global.InitDb(global.C.Dblink)
	Router := router.Routers()

	address := fmt.Sprintf(":%d", global.C.Addr)
	s := initServer(address, Router)

	global.LOG.Info("server run success on ", zap.String("address", address))

	global.LOG.Error(s.ListenAndServe().Error())
	// cmd.RunPrompt()
}
